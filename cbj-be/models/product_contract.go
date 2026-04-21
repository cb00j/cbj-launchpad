// models/product_contract.go
package models

import (
	"time"
)

type ProductContract struct {
	ID                        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name                      string    `gorm:"column:name;type:varchar(80);not null" json:"name"`
	Description               string    `gorm:"column:description;type:longtext" json:"description"`
	Img                       string    `gorm:"column:img;type:varchar(500)" json:"img"`
	TwitterName               string    `gorm:"column:twitter_name;type:varchar(40)" json:"twitterName"`
	Status                    int       `gorm:"column:status;not null;default:0" json:"status"`
	Amount                    string    `gorm:"column:amount;type:varchar(40)" json:"amount"`
	SaleContractAddress       string    `gorm:"column:sale_contract_address;type:varchar(42)" json:"saleContractAddress"`
	TokenAddress              string    `gorm:"column:token_address;type:varchar(42)" json:"tokenAddress"`
	PaymentToken              string    `gorm:"column:payment_token;type:varchar(42)" json:"paymentToken"`
	Follower                  int       `gorm:"column:follower;not null;default:0" json:"follower"`
	Tge                       time.Time `gorm:"column:tge" json:"tge"`
	ProjectWebsite            string    `gorm:"column:project_website;type:varchar(500)" json:"projectWebsite"`
	AboutHtml                 string    `gorm:"column:about_html;type:varchar(500)" json:"aboutHtml"`
	RegistrationTimeStarts    time.Time `gorm:"column:registration_time_starts" json:"registrationTimeStarts"`
	RegistrationTimeEnds      time.Time `gorm:"column:registration_time_ends" json:"registrationTimeEnds"`
	SaleStart                 time.Time `gorm:"column:sale_start" json:"saleStart"`
	SaleEnd                   time.Time `gorm:"column:sale_end" json:"saleEnd"`
	MaxParticipation          string    `gorm:"column:max_participation;type:varchar(40)" json:"maxParticipation"`
	TokenPriceInPT            string    `gorm:"column:token_price_in_PT;type:varchar(40)" json:"tokenPriceInPT"`
	TotalTokensSold           string    `gorm:"column:total_tokens_sold;type:varchar(40)" json:"totalTokensSold"`
	AmountOfTokensToSell      string    `gorm:"column:amount_of_tokens_to_sell;type:varchar(60)" json:"amountOfTokensToSell"`
	TotalRaised               string    `gorm:"column:total_raised;type:varchar(60)" json:"totalRaised"`
	Symbol                    string    `gorm:"column:symbol;type:varchar(60)" json:"symbol"`
	Decimals                  int       `gorm:"column:decimals;default:8" json:"decimals"`
	UnlockTime                time.Time `gorm:"column:unlock_time" json:"unlockTime"`
	Medias                    string    `gorm:"column:medias;type:varchar(200)" json:"medias"`
	NumberOfRegistrants       int       `gorm:"column:number_of_registrants" json:"numberOfRegistrants"`
	Vesting                   string    `gorm:"column:vesting;type:varchar(40)" json:"vesting"`
	Tricker                   string    `gorm:"column:tricker;type:varchar(40)" json:"tricker"`
	TokenName                 string    `gorm:"column:token_name;type:varchar(20)" json:"tokenName"`
	Roi                       string    `gorm:"column:roi;type:varchar(20)" json:"roi"`
	VestingPortionsUnlockTime string    `gorm:"column:vesting_portions_unlock_time;type:varchar(60)" json:"vestingPortionsUnlockTime"`
	VestingPercentPerPortion  string    `gorm:"column:vesting_percent_per_portion;type:varchar(60)" json:"vestingPercentPerPortion"`
	CreateTime                time.Time `gorm:"column:create_time;not null;autoCreateTime" json:"createTime"`
	UpdateTime                time.Time `gorm:"column:update_time;not null;autoUpdateTime" json:"updateTime"`
	Type                      int       `gorm:"column:type" json:"type"`
	CardLink                  string    `gorm:"column:card_link;type:varchar(200)" json:"cardLink"`
	TweetId                   string    `gorm:"column:tweet_id;type:varchar(40)" json:"tweetId"`
	ChainId                   int       `gorm:"column:chain_id;default:0" json:"chainId"`
	PaymentTokenDecimals      int       `gorm:"column:payment_token_decimals;default:8" json:"paymentTokenDecimals"`
	CurrentPrice              int64     `gorm:"column:current_price;default:0" json:"currentPrice"`
}

// TableName 指定表名
func (ProductContract) TableName() string {
	return "product_contract"
}
