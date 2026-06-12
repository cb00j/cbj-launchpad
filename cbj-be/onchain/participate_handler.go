package onchain

import (
	"cbj-be/contracts"
	"cbj-be/models"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type ParticipateHandler struct {
	client        *ethclient.Client
	db            *gorm.DB
	saleToProduct map[common.Address]uint64
	sig           common.Hash
}

func NewParticipateHandler(l *Listener) *ParticipateHandler {
	return &ParticipateHandler{
		client:        l.Client(),
		db:            l.DB(),
		saleToProduct: l.SaleToProduct(),
		sig:           l.ABI().Events["TokensSold"].ID,
	}
}

func (h *ParticipateHandler) EventSig() common.Hash {
	return h.sig
}

func (h *ParticipateHandler) Handle(vlog types.Log, filterer *contracts.CBJSaleFilterer) {
	event, err := filterer.ParseTokensSold(vlog)
	if err != nil {
		log.Printf("parse TokensSold failed: %v", err)
		return
	}

	saleAddress := event.Raw.Address
	productId, ok := h.saleToProduct[saleAddress]
	if !ok {
		log.Printf("saleAddress %s not found in saleToProduct map", saleAddress)
		return
	}

	userAddr := strings.ToLower(event.User.Hex())
	txHash := event.Raw.TxHash.Hex()
	blockNumber := event.Raw.BlockNumber

	log.Printf("TokensSold: user=%s amount=%s product=%d tx=%s",
		userAddr, event.Amount.String(), productId, txHash)

	// sync sale info to db
	h.syncSaleInfo(saleAddress, productId)
	// save participation info to db
	h.saveParticipation(saleAddress, productId, userAddr, txHash, blockNumber)

}

func (h *ParticipateHandler) syncSaleInfo(saleAddress common.Address, productId uint64) {
	sale, err := contracts.NewCBJSaleCaller(saleAddress, h.client)
	if err != nil {
		log.Printf("new sale caller failed: %v", err)
		return
	}

	result, err := sale.GetSaleInfo(nil)
	if err != nil {
		log.Printf("getSaleInfo failed: %v", err)
		return
	}

	err = h.db.Model(&models.ProductContract{}).Where("id = ?", productId).Updates(map[string]interface{}{
		"total_tokens_sold":        result.TotalTokensSold.String(),
		"total_raised":             result.TotalETHRaised.String(),
		"amount_of_tokens_to_sell": result.AmountOfTokensToSell.String(),
	}).Error

	if err != nil {
		log.Printf("update sale info failed: %v", err)
		return
	}
	log.Printf("synced sale info: product=%d sold=%s raised=%s",
		productId, result.TotalTokensSold.String(), result.TotalETHRaised.String())
}

func (h *ParticipateHandler) saveParticipation(saleAddress common.Address, productId uint64, userAddr, txHash string, blockNum uint64) {
	sale, err := contracts.NewCBJSaleCaller(saleAddress, h.client)
	if err != nil {
		log.Printf("new sale caller failed: %v", err)
		return
	}

	amountBought, amountETHPaid, timeParticipated, _, err := sale.GetParticipation(nil, common.HexToAddress(userAddr))
	if err != nil {
		log.Printf("getParticipation failed: %v", err)
		return
	}

	participation := models.ProductParticipation{
		ProductID:      productId,
		AccountID:      userAddr,
		AmountBought:   amountBought.String(),
		AmountPaid:     amountETHPaid.String(),
		TxHash:         txHash,
		BlockNumber:    blockNum,
		ParticipatedAt: time.Unix(timeParticipated.Int64(), 0),
		Status:         1,
	}

	result := h.db.Where("account_id = ? AND product_id = ?", userAddr, productId).FirstOrCreate(&participation)
	if result.Error != nil {
		log.Printf("save participation failed: %v", result.Error)
		return
	}
}
