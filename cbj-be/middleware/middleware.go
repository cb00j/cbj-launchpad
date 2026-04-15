package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Log(c *gin.Context) {
	path := c.Request.URL.Path
	username := c.GetString("username")
	fmt.Println("我是一条日志:", path, "用户:", username)

	context := c.Copy() // 重要:在goroutine中使用gin.Context需要先Copy创建深拷贝副本，因为gin.Context不是线程安全的且和主协程的生命周期不一致
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(context.GetString("username"))
	}()

	// 或者不要使用Context，使用传值的方式
	go func(uname, p string) {
		time.Sleep(1 * time.Second)
		fmt.Printf("用户: %s, 路径: %s\n", uname, p)
	}(username, path) // 传值，而不是传context
}

func Auth(c *gin.Context) {
	c.Set("username", "cbj") //中间件上下文设置值
	fmt.Println("权限校验成功")
}
