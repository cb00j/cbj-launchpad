package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (con BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func (con BaseController) Fail(c *gin.Context) {
	c.String(http.StatusInternalServerError, "fail")
}

func (con BaseController) FailMsg(c *gin.Context, message string) {
	c.String(http.StatusInternalServerError, message)
}

func (con BaseController) FailFormat(c *gin.Context, format string, data interface{}) {
	c.String(http.StatusInternalServerError, format, data)
}
