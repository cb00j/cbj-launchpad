package models

// ProductContract 产品合约信息
type ProductContract struct {
	// id
	ID int `gorm:"primaryKey;autoIncrement;comment:id" json:"id"`
	// 项目名称
	Name string `gorm:"type:varchar(200);comment:项目名称" json:"name"`
	// 项目描述
	Description string `gorm:"type:text;comment:项目描述" json:"description"`
	// 项目图标地址
	Img string `gorm:"type:varchar(500);comment:项目图标地址" json:"img"`
	// 项目状态
	Status int `gorm:"type:int;default:0;comment:项目状态" json:"status"`
	// 当前币种质押个数
	Amount string `gorm:"type:varchar(100);default:0;comment:当前币种质押个数" json:"amount"`
	// 已提取
	Withdrawed string `gorm:"-" json:"withdrawed"`
	// sale合约地址
	SaleContractAddress string `gorm:"type:varchar(255);comment:sale合约地址" json:"saleContractAddress"`
	// bre合约地址
	TokenAddress string `gorm:"type:varchar(255);comment:token合约地址" json:"tokenAddress"`
	// 支付代币地址
	PaymentToken string `gorm:"type:varchar(255);comment:支付代币地址" json:"paymentToken"`
	// ins或推特的follow数
	Follower int `gorm:"type:int;default:0;comment:关注人数" json:"follower"`
	// tge时间
	TGE int64 `gorm:"type:bigint;default:0;comment:TGE时间" json:"tge"`
	// 项目官网
	ProjectWebsite string `gorm:"type:varchar(500);comment:项目官网" json:"projectWebsite"`
	// 关于HTML
	AboutHtml string `gorm:"type:text;comment:关于HTML" json:"aboutHtml"`
	// 注册开始时间
	RegistrationTimeStarts int64 `gorm:"type:bigint;default:0;comment:注册开始时间" json:"registrationTimeStarts"`
	// 注册结束时间
	RegistrationTimeEnds int64 `gorm:"type:bigint;default:0;comment:注册结束时间" json:"registrationTimeEnds"`
	// sale开始时间
	SaleStart int64 `gorm:"type:bigint;default:0;comment:sale开始时间" json:"saleStart"`
	// sale结束时间
	SaleEnd int64 `gorm:"type:bigint;default:0;comment:sale结束时间" json:"saleEnd"`
	// 硬顶
	MaxParticipation string `gorm:"type:varchar(100);default:0;comment:硬顶" json:"maxParticipation"`
	// Token价格
	TokenPriceInPT string `gorm:"type:varchar(100);default:0;comment:Token价格" json:"tokenPriceInPT"`
	// 所有已卖的token个数
	TotalTokensSold string `gorm:"type:varchar(100);default:0;comment:所有已卖的token个数" json:"totalTokensSold"`
	// 待售代币数量
	AmountOfTokensToSell string `gorm:"type:varchar(100);default:0;comment:待售代币数量" json:"amountOfTokensToSell"`
	// 总募集金额
	TotalRaised string `gorm:"type:varchar(100);default:0;comment:总募集金额" json:"totalRaised"`
	// 代币符号
	Symbol string `gorm:"type:varchar(50);comment:代币符号" json:"symbol"`
	// 小数位数
	Decimals int `gorm:"type:int;default:18;comment:小数位数" json:"decimals"`
	// 解锁时间
	UnlockTime int64 `gorm:"type:bigint;default:0;comment:解锁时间" json:"unlockTime"`
	// 媒体链接
	Medias string `gorm:"type:text;comment:媒体链接" json:"medias"`
	// 注册人数
	NumberOfRegistrants int `gorm:"type:int;default:0;comment:注册人数" json:"numberOfRegistrants"`
	// 解锁计划
	Vesting string `gorm:"type:text;comment:解锁计划" json:"vesting"`
	// Tricker
	Tricker string `gorm:"type:varchar(100);comment:Tricker" json:"tricker"`
	// 代币名称
	TokenName string `gorm:"type:varchar(200);comment:代币名称" json:"tokenName"`
	// ROI
	ROI string `gorm:"type:varchar(100);default:0;comment:ROI" json:"roi"`
	// 解锁时间点列表
	VestingPortionsUnlockTime string `gorm:"type:text;comment:解锁时间点列表" json:"vestingPortionsUnlockTime"`
	// 每期解锁百分比
	VestingPercentPerPortion string `gorm:"type:text;comment:每期解锁百分比" json:"vestingPercentPerPortion"`
	// 创建时间
	CreateTime int64 `gorm:"type:bigint;default:0;comment:创建时间" json:"createTime"`
	// 更新时间
	UpdateTime int64 `gorm:"type:bigint;default:0;comment:更新时间" json:"updateTime"`
	// 类型
	Type int `gorm:"type:int;default:0;comment:类型" json:"type"`
	// 卡片链接
	CardLink string `gorm:"type:varchar(500);comment:卡片链接" json:"cardLink"`
	// 转推任务的推文ID
	TweetId string `gorm:"type:varchar(100);comment:转推任务的推文ID" json:"tweetId"`
	// 链ID
	ChainId int `gorm:"type:int;default:0;comment:链ID" json:"chainId"`
	// 支付代币小数位数
	PaymentTokenDecimals int `gorm:"type:int;default:18;comment:支付代币小数位数" json:"paymentTokenDecimals"`
	// 当前价格
	CurrentPrice string `gorm:"type:varchar(100);default:0;comment:当前价格" json:"currentPrice"`
}

// TableName 指定表名
func (ProductContract) TableName() string {
	return "product_contract"
}
