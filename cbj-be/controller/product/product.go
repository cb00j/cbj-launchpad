package product

import (
	"cbj-be/controller/base"
	"cbj-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	base.BaseController
}

func (con ProductController) List(c *gin.Context) {
	var productContractList []models.ProductContract
	models.DB.Find(&productContractList)
	c.JSON(http.StatusOK, gin.H{
		"data": productContractList,
	})
}
