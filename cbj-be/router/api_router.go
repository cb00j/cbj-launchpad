package router

import (
	"cbj-be/controller/encode"
	"cbj-be/controller/product"
	"cbj-be/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.RouterGroup, productController *product.ProductController, encodeController *encode.EncodeController, allocationController *product.AllocationController, registerController *product.RegisterController) {
	apiRouter := r.Group("api", middleware.Auth, middleware.Log)
	apiRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是前台首页",
		})
	})

	productRouter := apiRouter.Group("product")
	productRouter.GET("/list", productController.List)
	productRouter.GET("/base_info", productController.BaseInfo)
	productRouter.GET("/apr", productController.Apr)
	productRouter.POST("/allocation/calc", allocationController.Calc)
	productRouter.POST("/register/user_register", registerController.UserRegister)

	encodeRouter := apiRouter.Group("encode")
	encodeRouter.POST("/sign_registration", encodeController.SignRegistration)
	encodeRouter.POST("/sign_participation", encodeController.SignParticipation)
}
