package product

import (
	"cbj-be/models"
	"cbj-be/utils"
	"cbj-be/vo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
}

func (con ProductController) BaseInfo(c *gin.Context) {
	productId, err := utils.Int(c.Query("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid product ID",
		})
		return
	}
	productContract := models.ProductContract{ID: uint(productId)}
	result := con.db.First(&productContract)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    500,
			"message": result.Error.Error(),
		})
		return
	}

	vo := vo.FromModel(&productContract)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    vo,
		"message": "success",
	})
}

func (con ProductController) List(c *gin.Context) {
	var productContractList []models.ProductContract
	con.db.Find(&productContractList)

	voList := make([]vo.ProductContractVO, 0, len(productContractList))
	for _, product := range productContractList {
		voList = append(voList, *vo.FromModel(&product))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    voList,
		"message": "success",
	})
}

func (con ProductController) Apr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"apr":       0.89,
		"priceInLP": 0.5,
		"message":   "success",
	})
}
