package router

import (
	"cbj-be/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.RouterGroup) {
	adminRouter := r.Group("admin", middleware.Auth, middleware.Log)
	adminRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "我是后台首页",
		})
	})
}
