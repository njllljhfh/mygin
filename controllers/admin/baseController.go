package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BaseController 用于其他的Controller结构体继承
type BaseController struct{}

func (con *BaseController) success(c *gin.Context, data interface{}) {
	c.String(http.StatusOK, "成功: %v", data)
}

func (con *BaseController) error(c *gin.Context, data interface{}) {
	c.String(http.StatusOK, "失败: %v", data)
}
