package routers

import (
	"github.com/gin-gonic/gin"
	"mygin/controllers/api"
)

func ApiRouterInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		indexController := &api.IndexController{}
		apiRouters.GET("/", indexController.Index)

		ectController := &api.EtcController{}
		apiRouters.GET("/userlist", ectController.UserList)
		apiRouters.GET("/plist", ectController.Plist)
	}
}
