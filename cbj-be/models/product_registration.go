package models

import (
	"cbj-be/vo"
	"time"
)

type ProductRegistration struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AccountID    string    `gorm:"column:account_id;not null" json:"accountId"`       // 用户钱包地址
	ProductID    uint64    `gorm:"column:product_id;not null" json:"productId"`       // 项目ID
	ReferralCode string    `gorm:"column:referral_code" json:"referralCode"`          // 推荐码(可空)
	StakeAmount  string    `gorm:"column:stake_amount;default:0" json:"stakeAmount"`  // 质押量(wei,字符串存大数)
	Allocation   string    `gorm:"column:allocation;default:0" json:"allocation"`     // 认购额度(wei)
	TxHash       string    `gorm:"column:tx_hash" json:"txHash"`                      // 链上交易哈希
	Status       int8      `gorm:"column:status;default:1" json:"status"`             // 状态:1-已注册,0-无效
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"` // 注册时间
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"` // 更新时间
}

// TableName 指定表名
func (ProductRegistration) TableName() string {
	return "product_registration"
}

func (p ProductRegistration) ToVO() *vo.ProductRegistrationVO {
	return &vo.ProductRegistrationVO{
		ID:           p.ID,
		AccountID:    p.AccountID,
		ProductID:    p.ProductID,
		ReferralCode: p.ReferralCode,
		StakeAmount:  p.StakeAmount,
		Allocation:   p.Allocation,
		TxHash:       p.TxHash,
		Status:       p.Status,
		CreatedAt:    p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
