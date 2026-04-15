package router

import (
	"cbj-be/controller/product"
	"cbj-be/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	apiRouter := r.Group("api", middleware.Auth, middleware.Log)
	apiRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是前台首页",
		})
	})

	apiRouter.GET("boba/product/list", product.ProductController{}.List)
}
