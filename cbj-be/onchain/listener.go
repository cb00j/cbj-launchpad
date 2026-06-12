package onchain

import (
	"cbj-be/contracts"
	"cbj-be/models"
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const syncStateName = "sale_listener"
const pollBatchSize = uint64(2000)

type EventHandler interface {
	EventSig() common.Hash
	Handle(vlog types.Log, filterer *contracts.CBJSaleFilterer)
}

type Listener struct {
	client        *ethclient.Client
	db            *gorm.DB
	abi           abi.ABI
	saleToProduct map[common.Address]uint64 // sale address to product id
	handlers      map[common.Hash]EventHandler
}

func NewListener(wsUrl string, db *gorm.DB) (*Listener, error) {
	client, err := ethclient.Dial(wsUrl)
	if err != nil {
		return nil, err
	}
	//parsedABI, err := abi.JSON(strings.NewReader(cbjSaleABIJson))
	parsedABI, err := abi.JSON(strings.NewReader(contracts.CBJSaleMetaData.ABI))
	if err != nil {
		return nil, err
	}

	return &Listener{
		client:        client,
		db:            db,
		abi:           parsedABI,
		saleToProduct: make(map[common.Address]uint64),
		handlers:      make(map[common.Hash]EventHandler),
	}, nil
}

func (l *Listener) ABI() abi.ABI {
	return l.abi
}

func (l *Listener) Client() *ethclient.Client {
	return l.client
}

func (l *Listener) DB() *gorm.DB {
	return l.db
}

func (l *Listener) SaleToProduct() map[common.Address]uint64 {
	return l.saleToProduct
}

func (l *Listener) RegisterHandler(handler EventHandler) {
	l.handlers[handler.EventSig()] = handler
}

// register a sale contract to listener
func (l *Listener) RegisterSale(saleAddress string, productId uint64) {
	l.saleToProduct[common.HexToAddress(saleAddress)] = productId
}

func (l *Listener) LoadSalesFromDB() error {
	var productContracts []models.ProductContract
	if err := l.db.Model(&models.ProductContract{}).Find(&productContracts).Error; err != nil {
		return err
	}

	for _, productContract := range productContracts {
		l.RegisterSale(productContract.SaleContractAddress, uint64(productContract.ID))
	}

	return nil
}

func (l *Listener) allEventSigs() []common.Hash {
	eventSigs := make([]common.Hash, 0, len(l.handlers))
	for _, handler := range l.handlers {
		eventSigs = append(eventSigs, handler.EventSig())
	}
	return eventSigs
}

func (l *Listener) allAddresses() []common.Address {
	addresses := make([]common.Address, 0, len(l.saleToProduct))
	for address := range l.saleToProduct {
		addresses = append(addresses, address)
	}
	return addresses
}

func (l *Listener) Start(ctx context.Context) {
	for {
		// poll missed blocks first
		l.pollMissed(ctx)
		// then watch for new events
		err := l.watch(ctx)
		if err != nil {
			log.Printf("listener watch error, reconnect in 5s: %v", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
				continue
			}
		}
		return
	}
}

func (l *Listener) pollMissed(ctx context.Context) {
	fromBlock, err := l.getLastBlock()
	if err != nil {
		log.Printf("pollMissed get last block failed: %v", err)
		return
	}

	latest, err := l.client.BlockNumber(ctx)
	if err != nil {
		log.Printf("pollMissed get latest block failed: %v", err)
		return
	}

	if fromBlock >= latest {
		log.Printf("no missed blocks to poll")
		return // no missed blocks
	}

	addresses := l.allAddresses()
	if len(addresses) == 0 {
		return
	}

	log.Printf("pollMissed: scanning blocks %d -> %d", fromBlock, latest)
	// used to parse log
	filterer, err := contracts.NewCBJSaleFilterer(common.Address{}, nil)
	if err != nil {
		log.Printf("pollMissed new filterer failed: %v", err)
		return
	}

	for start := fromBlock; start < latest; start += pollBatchSize {
		end := start + pollBatchSize - 1
		if end > latest {
			end = latest
		}
		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: addresses,
			Topics:    [][]common.Hash{l.allEventSigs()},
		}

		logs, err := l.client.FilterLogs(ctx, query)
		if err != nil {
			log.Printf("pollMissed FilterLogs %d-%d failed: %v", start, end, err)
			return
		}
		for _, vlog := range logs {
			l.dispatch(vlog, filterer)
		}
		l.saveLastBlock(end) // save the last block number polled

	}
	log.Printf("pollMissed: done up to block %d", latest)

}

func (l *Listener) watch(ctx context.Context) error {
	addresses := l.allAddresses()
	if len(addresses) == 0 {
		log.Println("no sale contracts to watch")
		return nil
	}

	query := ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    [][]common.Hash{l.allEventSigs()},
	}

	logs := make(chan types.Log)
	sub, err := l.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	filterer, err := contracts.NewCBJSaleFilterer(common.Address{}, nil)
	if err != nil {
		return err
	}

	log.Printf("listening %d events on %d contracts", len(l.handlers), len(addresses))

	for {
		select {
		case err := <-sub.Err():
			return err
		case vlog := <-logs:
			l.dispatch(vlog, filterer)
			l.saveLastBlock(vlog.BlockNumber)
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}

func (l *Listener) dispatch(vlog types.Log, filterer *contracts.CBJSaleFilterer) {
	if len(l.handlers) == 0 {
		return
	}

	if h, ok := l.handlers[vlog.Topics[0]]; ok {
		h.Handle(vlog, filterer)
	}
}

func (l *Listener) getLastBlock() (uint64, error) {
	var state models.SyncState
	err := l.db.Where("name = ?", syncStateName).First(&state).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil // first time, no state, return 0
	}
	if err != nil {
		return 0, err
	}
	return state.LastBlock, nil
}

func (l *Listener) saveLastBlock(block uint64) {
	state := models.SyncState{
		Name:      syncStateName,
		LastBlock: block,
	}

	// update if record exists, insert if not
	err := l.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_block", "updated_at"}),
	}).Create(&state).Error
	if err != nil {
		log.Printf("saveLastBlock failed: %v", err)
	}
}
