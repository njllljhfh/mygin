package routers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "mygin/controllers/itying"
)

func initMiddleware(c *gin.Context) {
    fmt.Println("中间件: 执行 处理函数 前 --- 1")
    // 调用该请求的剩余处理程序
    c.Next()
    fmt.Println("中间件: 执行 处理函数 前 --- 2")
}

/*
web 的 index 接口
*/

func DefaultRouterInit(r *gin.Engine) {
    defaultRouters := r.Group("/")
    {
        indexController := &itying.IndexController{}
        defaultRouters.GET("/", initMiddleware, indexController.Index)

        newsController := &itying.NewsController{}
        defaultRouters.GET("/news", initMiddleware, newsController.News)
    }
}
