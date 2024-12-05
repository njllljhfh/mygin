package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsController struct{}

func (con *NewsController) News(c *gin.Context) {
	// 第二个参数是 html中define定义的名字
	c.HTML(http.StatusOK, "adminNews", gin.H{})
}
