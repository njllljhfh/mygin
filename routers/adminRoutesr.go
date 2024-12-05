package routers

import (
	"github.com/gin-gonic/gin"
	"mygin/controllers/admin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		indexController := &admin.IndexController{}
		adminRouters.GET("/", indexController.Index)

		newsController := &admin.NewsController{}
		adminRouters.GET("/news", newsController.News)

		userController := &admin.UserController{}
		adminRouters.GET("/user", userController.Index)
		adminRouters.GET("/user/add", userController.Add)
		adminRouters.GET("/user/edit", userController.Edit)

		articleController := &admin.ArticleController{}
		adminRouters.GET("/article", articleController.Index)
		adminRouters.GET("/article/add", articleController.Add)
		adminRouters.GET("/article/edit", articleController.Edit)
	}
}
