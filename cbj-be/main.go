package main

import (
	"cbj-be/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	router.ApiRouterInit(r)
	http.ListenAndServe(":8080", r)
	r.Run()
}
