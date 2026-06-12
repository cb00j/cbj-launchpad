package models

import "time"

// ProductParticipation 用户项目认购明细表,对应链上 TokensSold 事件
type ProductParticipation struct {
	ID             uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AccountID      string    `gorm:"column:account_id;not null" json:"accountId"`
	ProductID      uint64    `gorm:"column:product_id;not null" json:"productId"`
	AmountBought   string    `gorm:"column:amount_bought;default:0" json:"amountBought"` // 购买代币数量(wei)
	AmountPaid     string    `gorm:"column:amount_paid;default:0" json:"amountPaid"`     // 支付ETH(wei)
	TxHash         string    `gorm:"column:tx_hash" json:"txHash"`
	BlockNumber    uint64    `gorm:"column:block_number;default:0" json:"blockNumber"`
	ParticipatedAt time.Time `gorm:"column:participated_at" json:"participatedAt"`
	Status         int8      `gorm:"column:status;default:1" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (ProductParticipation) TableName() string {
	return "product_participation"
}
