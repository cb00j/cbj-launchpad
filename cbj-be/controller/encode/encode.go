package encode

import (
	"cbj-be/controller/base"
	"cbj-be/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type EncodeController struct {
	base.BaseController
}

func (con EncodeController) SignRegistration(c *gin.Context) {
	// 1. get parameters
	userAddress := c.PostForm("userAddress")
	contractAddress := c.PostForm("contractAddress")

	// 2. validate parameters
	if strings.TrimSpace(userAddress) == "" || strings.TrimSpace(contractAddress) == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "userAddress and contractAddress are required",
		})
		return
	}

	// 3.handle 0x prefix and case sensitivity, make them all lowercase and remove 0x prefix if exists
	userAddress = strings.TrimPrefix(strings.ToLower(userAddress), "0x")
	contractAddress = strings.TrimPrefix(strings.ToLower(contractAddress), "0x")
	concat := userAddress + contractAddress
	hexStr := "0x" + concat

	// 4.sign by private key and return signature
	sign, err := utils.Credentials.GetSign(hexStr)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "failed to get sign: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    sign,
		"message": "success",
	})

}

func (con EncodeController) SignParticipation(c *gin.Context) {

}
