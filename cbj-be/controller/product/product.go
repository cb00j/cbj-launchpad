package product

import (
	"cbj-be/controller/base"
	"cbj-be/models"
	"cbj-be/utils"
	"cbj-be/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	base.BaseController
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
	result := models.DB.First(&productContract)
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
	models.DB.Find(&productContractList)

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
