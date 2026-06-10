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

const syncStateName = "register_listener"

type RegisterListener struct {
	client        *ethclient.Client
	db            *gorm.DB
	eventSig      common.Hash
	saleToProduct map[common.Address]uint64
}

func NewRegisterListener(wsURL string, db *gorm.DB) (*RegisterListener, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	//parsedABI, err := abi.JSON(strings.NewReader(cbjSaleABIJson))
	parsedABI, err := abi.JSON(strings.NewReader(contracts.CBJSaleMetaData.ABI))
	if err != nil {
		return nil, err
	}

	// event UserRegistered signature hash
	//eventSig := crypto.Keccak256Hash([]byte("UserRegistered(address,uint256)"))
	eventSig := parsedABI.Events["UserRegistered"].ID

	return &RegisterListener{
		client:        client,
		db:            db,
		eventSig:      eventSig,
		saleToProduct: make(map[common.Address]uint64),
	}, nil
}

func (l *RegisterListener) RegisterSale(saleAddress string, productID uint64) {
	l.saleToProduct[common.HexToAddress(saleAddress)] = productID
}

func (l *RegisterListener) Watch(ctx context.Context) error {
	addresses := make([]common.Address, 0, len(l.saleToProduct))
	for saleAddress := range l.saleToProduct {
		addresses = append(addresses, saleAddress)
	}

	query := ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    [][]common.Hash{{l.eventSig}},
	}

	logs := make(chan types.Log)
	sub, err := l.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}
	log.Println("subscription established, watching...")
	defer sub.Unsubscribe()

	// used to parse log
	filterer, err := contracts.NewCBJSaleFilterer(common.Address{}, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vlog := <-logs:
			event, err := filterer.ParseUserRegistered(vlog)
			if err != nil {
				log.Printf("ParseUserRegistered failed: %v", err)
				continue
			}
			l.handleEvent(event)
			l.saveLastBlock(vlog.BlockNumber)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (l *RegisterListener) handleEvent(event *contracts.CBJSaleUserRegistered) {
	saleAddress := event.Raw.Address
	productId, ok := l.saleToProduct[saleAddress]
	if !ok {
		return // not concerned about this sale
	}

	userAddr := strings.ToLower(event.User.Hex())
	txHash := event.Raw.TxHash.Hex()
	blockNum := event.Raw.BlockNumber // block number of the transaction

	log.Printf("User %s registered product %d in sale %s at block %d", userAddr, productId, userAddr, blockNum)

	l.saveRegistration(userAddr, productId, txHash)
}

func (l *RegisterListener) saveRegistration(userAddr string, productId uint64, txHash string) {
	reg := models.ProductRegistration{
		AccountID: userAddr,
		ProductID: productId,
		TxHash:    txHash,
		Status:    1,
	}
	result := l.db.Where("account_id = ? AND product_id = ?", userAddr, productId).
		FirstOrCreate(&reg)
	if result.Error != nil {
		log.Printf("save registration failed: %v", result.Error)
	}
}

func (l *RegisterListener) Start(ctx context.Context) {
	for {
		// poll missed blocks first
		l.pollMissed(ctx)
		// then watch for new events
		err := l.Watch(ctx)
		if err != nil {
			log.Printf("watch error, reconnect in 5s: %v", err)
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

func (l *RegisterListener) pollMissed(ctx context.Context) {
	fromBlock, err := l.getLastBlock()
	if err != nil {
		log.Printf("pollMissed: get last block failed: %v", err)
		return
	}

	lastest, err := l.client.BlockNumber(ctx)
	if err != nil {
		log.Printf("pollMissed: get latest block failed: %v", err)
		return
	}

	if fromBlock >= lastest {
		return // no missed blocks
	}

	log.Printf("pollMissed: scanning blocks from %d to %d", fromBlock, lastest)

	addresses := make([]common.Address, 0, len(l.saleToProduct))
	for saleAddress := range l.saleToProduct {
		addresses = append(addresses, saleAddress)
	}

	if len(addresses) == 0 {
		return
	}

	// batch query
	const batchSize = uint64(2000)
	filterer, err := contracts.NewCBJSaleFilterer(common.Address{}, nil)
	if err != nil {
		log.Printf("pollMissed: create filterer failed: %v", err)
		return
	}

	for start := fromBlock + 1; start <= lastest; start += batchSize {
		end := start + batchSize - 1
		if end > lastest {
			end = lastest
		}
		query := ethereum.FilterQuery{
			Addresses: addresses,
			Topics:    [][]common.Hash{{l.eventSig}},
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
		}

		logs, err := l.client.FilterLogs(ctx, query)
		if err != nil {
			log.Printf("pollMissed: get logs failed: %v", err)
			return
		}

		for _, vlog := range logs {
			event, err := filterer.ParseUserRegistered(vlog)
			if err != nil {
				log.Printf("pollMissed: ParseUserRegistered failed: %v", err)
				continue
			}
			l.handleEvent(event)
		}

		if err := l.saveLastBlock(end); err != nil {
			log.Printf("pollMissed: save last block failed: %v", err)
		}
	}

	log.Printf("pollMissed: done,processed up to block %d", lastest)
}

func (l *RegisterListener) getLastBlock() (uint64, error) {
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

func (l *RegisterListener) saveLastBlock(lastBlock uint64) error {
	state := models.SyncState{Name: syncStateName, LastBlock: lastBlock}

	return l.db.Clauses(clause.OnConflict{
		// update if record exists, insert if not
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"last_block", "updated_at"}),
	}).Create(&state).Error
}

func (l *RegisterListener) LoadSalesFromDB() error {
	ProductContracts := []models.ProductContract{}
	err := l.db.Find(&ProductContracts).Error
	if err != nil {
		return err
	}

	for _, productContract := range ProductContracts {
		saleAddress := productContract.SaleContractAddress
		productID := productContract.ID
		l.RegisterSale(saleAddress, uint64(productID))
	}

	return nil
}
