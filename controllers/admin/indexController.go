package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

// Index 后台首页
func (con *IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}
