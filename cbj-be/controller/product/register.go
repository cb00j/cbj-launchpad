package product

import (
	"cbj-be/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterController struct {
	db *gorm.DB
}

func NewRegisterController(db *gorm.DB) *RegisterController {
	return &RegisterController{db: db}
}

func (con *RegisterController) UserRegister(c *gin.Context) {
	accountId := c.PostForm("accountId")
	productId := c.PostForm("productId")
	referralCode := c.PostForm("referralCode")
	txHash := c.PostForm("txHash")

	// Validate input
	if strings.TrimSpace(accountId) == "" || strings.TrimSpace(productId) == "" || strings.TrimSpace(txHash) == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	accountId = strings.ToLower(accountId)
	parsedProductID, err := strconv.ParseUint(productId, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	pr := &models.ProductRegistration{
		AccountID:    accountId,
		ProductID:    parsedProductID,
		ReferralCode: referralCode,
		TxHash:       txHash,
	}

	// Save to database
	if err := con.db.Create(pr).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})

}
