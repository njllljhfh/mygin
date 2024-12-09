package routers

import (
    "github.com/gin-gonic/gin"
    "mygin/controllers/admin"
    MDW "mygin/middlewares"
)

func AdminRouterInit(r *gin.Engine) {
    adminRouters := r.Group("/admin")
    // 在当前的分组路由中，配置中间件，方法1
    // adminRouters := r.Group("/admin", MDW.AdminMiddlewareOne, MDW.AdminMiddlewareTwo)
    // 在当前的分组路由中，配置中间件，这些中间件只会在访问 /admin 下的接口时才会被调用
    // 在当前的分组路由中，配置中间件，方法2
    adminRouters.Use(
        MDW.AdminMiddlewareOne,
        MDW.AdminMiddlewareTwo,
    )
    {
        indexController := &admin.IndexController{}
        adminRouters.GET("/", indexController.Index)

        newsController := &admin.NewsController{}
        adminRouters.GET("/news", newsController.News)

        userController := &admin.UserController{}
        adminRouters.GET("/user", userController.User) // 获取中间件中添加的数据
        adminRouters.POST("/user/add", userController.UserAdd)
        adminRouters.GET("/user/edit", userController.UserEdit)

        articleController := &admin.ArticleController{}
        adminRouters.GET("/article", articleController.Article)
        adminRouters.GET("/article/add", articleController.ArticleAdd)
        adminRouters.GET("/article/edit", articleController.ArticleEdit)
    }
}
