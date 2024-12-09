package routers

import (
    "github.com/gin-gonic/gin"
    "mygin/controllers/itying"
    MDW "mygin/middlewares"
)

/*
web 的 index 接口
*/

func DefaultRouterInit(r *gin.Engine) {
    defaultRouters := r.Group("/")
    {
        indexController := &itying.IndexController{}
        defaultRouters.GET("/", MDW.InitMiddlewareOne, MDW.InitMiddlewareTwo, indexController.Index)

        newsController := &itying.NewsController{}
        defaultRouters.GET("/news", MDW.InitMiddlewareOne, newsController.News)
    }
}
