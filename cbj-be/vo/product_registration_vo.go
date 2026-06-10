package vo

type ProductRegistrationVO struct {
	ID           uint64 `json:"id"`
	AccountID    string `json:"accountId"`    // 用户钱包地址
	ProductID    uint64 `json:"productId"`    // 项目ID
	ReferralCode string `json:"referralCode"` // 推荐码
	StakeAmount  string `json:"stakeAmount"`  // 质押量(wei)
	Allocation   string `json:"allocation"`   // 认购额度(wei)
	TxHash       string `json:"txHash"`       // 链上交易哈希
	Status       int8   `json:"status"`       // 状态
	CreatedAt    string `json:"createdAt"`    // 注册时间(格式化字符串)
	UpdatedAt    string `json:"updatedAt"`    // 更新时间
}
