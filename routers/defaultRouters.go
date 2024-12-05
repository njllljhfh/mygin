package routers

import (
	"github.com/gin-gonic/gin"
	"mygin/controllers/itying"
)

func DefaultRouterInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		indexController := &itying.IndexController{}
		defaultRouters.GET("/", indexController.Index)

		newsController := &itying.NewsController{}
		defaultRouters.GET("/news", newsController.News)
	}
}
