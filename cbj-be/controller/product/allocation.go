package product

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
)

type AllocationController struct {
}

func NewAllocationController() *AllocationController {
	return &AllocationController{}
}

func (con *AllocationController) Calc(c *gin.Context) {
	accountId := c.PostForm("accountId")
	productId := c.PostForm("productId")

	amount := new(big.Int).Mul(
		big.NewInt(1000),
		big.NewInt(params.Ether),
	)

	// connect to AllocationStaking contract and get the staking amount, then calculate the allocation amount
	c.JSON(200, gin.H{
		"accountId": accountId,
		"productId": productId,
		"amount":    amount.String(),
	})
}
