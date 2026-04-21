package router

import (
	"cbj-be/controller/product"
	"cbj-be/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.RouterGroup) {
	apiRouter := r.Group("api", middleware.Auth, middleware.Log)
	apiRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是前台首页",
		})
	})

	productRouter := apiRouter.Group("product")
	productRouter.GET("/list", product.ProductController{}.List)
	productRouter.GET("/base_info", product.ProductController{}.BaseInfo)
}
