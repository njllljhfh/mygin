package itying

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsController struct{}

// News 前台新闻
func (con *NewsController) News(c *gin.Context) {
	news := &Article{
		Title:   "新闻-标题",
		Desc:    "新闻-描述",
		Content: "新闻-详情",
	}
	// 返回html模板
	c.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "新闻页面",
		"news":  news,
	})
}
