// vo/product_contract_vo.go
package vo

import (
	"cbj-be/models"
	"encoding/json"
)

// 时间字段统一转为毫秒时间戳（Long）返回给前端
type ProductContractVO struct {
	ID                        uint    `json:"id"`
	Name                      string  `json:"name"`
	Description               string  `json:"description"`
	Img                       string  `json:"img"`
	TwitterName               string  `json:"twitterName"`
	Status                    int     `json:"status"`
	Amount                    string  `json:"amount"`
	SaleContractAddress       string  `json:"saleContractAddress"`
	TokenAddress              string  `json:"tokenAddress"`
	PaymentToken              string  `json:"paymentToken"`
	Follower                  int     `json:"follower"`
	Tge                       int64   `json:"tge"` // 时间戳（毫秒）
	ProjectWebsite            string  `json:"projectWebsite"`
	AboutHtml                 string  `json:"aboutHtml"`
	RegistrationTimeStarts    int64   `json:"registrationTimeStarts"`
	RegistrationTimeEnds      int64   `json:"registrationTimeEnds"`
	SaleStart                 int64   `json:"saleStart"`
	SaleEnd                   int64   `json:"saleEnd"`
	MaxParticipation          string  `json:"maxParticipation"`
	TokenPriceInPT            string  `json:"tokenPriceInPT"`
	TotalTokensSold           string  `json:"totalTokensSold"`
	AmountOfTokensToSell      string  `json:"amountOfTokensToSell"`
	TotalRaised               string  `json:"totalRaised"`
	Symbol                    string  `json:"symbol"`
	Decimals                  int     `json:"decimals"`
	UnlockTime                int64   `json:"unlockTime"`
	Medias                    string  `json:"medias"`
	NumberOfRegistrants       int     `json:"numberOfRegistrants"`
	Vesting                   string  `json:"vesting"`
	Tricker                   string  `json:"tricker"`
	TokenName                 string  `json:"tokenName"`
	Roi                       string  `json:"roi"`
	VestingPortionsUnlockTime []int64 `json:"vestingPortionsUnlockTime"` // 字符串转数组
	VestingPercentPerPortion  []int64 `json:"vestingPercentPerPortion"`
	CreateTime                int64   `json:"createTime"`
	UpdateTime                int64   `json:"updateTime"`
	Type                      int     `json:"type"`
	CardLink                  string  `json:"cardLink"`
	TweetId                   string  `json:"tweetId"`
	ChainId                   int     `json:"chainId"`
	PaymentTokenDecimals      int     `json:"paymentTokenDecimals"`
	CurrentPrice              int64   `json:"currentPrice"`
}

// FromModel 把 Model 转换为 VO
func FromModel(p *models.ProductContract) *ProductContractVO {
	vo := &ProductContractVO{
		ID:                     p.ID,
		Name:                   p.Name,
		Description:            p.Description,
		Img:                    p.Img,
		TwitterName:            p.TwitterName,
		Status:                 p.Status,
		Amount:                 p.Amount,
		SaleContractAddress:    p.SaleContractAddress,
		TokenAddress:           p.TokenAddress,
		PaymentToken:           p.PaymentToken,
		Follower:               p.Follower,
		Tge:                    p.Tge.UnixMilli(),
		ProjectWebsite:         p.ProjectWebsite,
		AboutHtml:              p.AboutHtml,
		RegistrationTimeStarts: p.RegistrationTimeStarts.UnixMilli(),
		RegistrationTimeEnds:   p.RegistrationTimeEnds.UnixMilli(),
		SaleStart:              p.SaleStart.UnixMilli(),
		SaleEnd:                p.SaleEnd.UnixMilli(),
		MaxParticipation:       p.MaxParticipation,
		TokenPriceInPT:         p.TokenPriceInPT,
		TotalTokensSold:        p.TotalTokensSold,
		AmountOfTokensToSell:   p.AmountOfTokensToSell,
		TotalRaised:            p.TotalRaised,
		Symbol:                 p.Symbol,
		Decimals:               p.Decimals,
		UnlockTime:             p.UnlockTime.UnixMilli(),
		Medias:                 p.Medias,
		NumberOfRegistrants:    p.NumberOfRegistrants,
		Vesting:                p.Vesting,
		Tricker:                p.Tricker,
		TokenName:              p.TokenName,
		Roi:                    p.Roi,
		CreateTime:             p.CreateTime.UnixMilli(),
		UpdateTime:             p.UpdateTime.UnixMilli(),
		Type:                   p.Type,
		CardLink:               p.CardLink,
		TweetId:                p.TweetId,
		ChainId:                p.ChainId,
		PaymentTokenDecimals:   p.PaymentTokenDecimals,
		CurrentPrice:           p.CurrentPrice,
	}

	// 把字符串转为 []int64（对应 Java 的 JSONArray.parseArray）
	if p.VestingPortionsUnlockTime != "" {
		_ = json.Unmarshal([]byte(p.VestingPortionsUnlockTime), &vo.VestingPortionsUnlockTime)
	}
	if p.VestingPercentPerPortion != "" {
		_ = json.Unmarshal([]byte(p.VestingPercentPerPortion), &vo.VestingPercentPerPortion)
	}

	return vo
}
