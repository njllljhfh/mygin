package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct{}

func (con *ArticleController) Index(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表")
}

func (con *ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-add")
}

func (con *ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "新闻列表-edit")
}
