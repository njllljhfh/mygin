package admin

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type ArticleController struct{}

func (con *ArticleController) Article(c *gin.Context) {
    c.String(http.StatusOK, "新闻列表")
}

func (con *ArticleController) ArticleAdd(c *gin.Context) {
    c.String(http.StatusOK, "新闻列表-add")
}

func (con *ArticleController) ArticleEdit(c *gin.Context) {
    c.String(http.StatusOK, "新闻列表-edit")
}
