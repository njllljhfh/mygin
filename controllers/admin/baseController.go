package admin

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// BaseController 用于其他的Controller结构体继承
type BaseController struct{}

func (con *BaseController) success(c *gin.Context, msg string, data interface{}) {
    // c.String(http.StatusOK, "成功: %v", data)
    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "msg":  "成功：" + msg,
        "data": data,
    })
}

func (con *BaseController) error(c *gin.Context, msg string, data interface{}) {
    // c.String(http.StatusOK, "失败: %v", data)
    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "msg":  "失败：" + msg,
        "data": data,
    })
}
