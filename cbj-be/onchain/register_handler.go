package onchain

import (
	"cbj-be/contracts"
	"cbj-be/models"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type RegisterHandler struct {
	db            *gorm.DB
	saleToProduct map[common.Address]uint64
	sig           common.Hash
}

func NewRegisterHandler(l *Listener) *RegisterHandler {
	return &RegisterHandler{
		db:            l.DB(),
		saleToProduct: l.SaleToProduct(),
		sig:           l.ABI().Events["UserRegister"].ID,
	}
}

func (h *RegisterHandler) EventSig() common.Hash {
	return h.sig
}

func (h *RegisterHandler) Handle(vlog types.Log, filterer *contracts.CBJSaleFilterer) {
	event, err := filterer.ParseUserRegistered(vlog)
	if err != nil {
		log.Printf("parse UserRegistered failed: %v", err)
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

	log.Printf("UserRegistered: user=%s product=%d tx=%s", userAddr, productId, txHash)

	reg := models.ProductRegistration{
		AccountID: userAddr,
		ProductID: productId,
		TxHash:    txHash,
		Status:    1,
	}

	result := h.db.Where("account_id = ? AND product_id = ?", userAddr, productId).FirstOrCreate(&reg)

	if result.Error != nil {
		log.Printf("save registration failed: %v", result.Error)
	}
}
