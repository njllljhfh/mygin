package routers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func AdminRouterInit(r *gin.Engine) {
    adminRouters := r.Group("/admin")
    {
        adminRouters.GET("/", func(c *gin.Context) {
            c.HTML(http.StatusOK, "admin/index.html", gin.H{})
        })
        adminRouters.GET("/news", func(c *gin.Context) {
            // 第二个参数是 html中define定义的名字
            c.HTML(http.StatusOK, "adminNews", gin.H{})
        })

        adminRouters.GET("/user", func(c *gin.Context) {
            c.String(http.StatusOK, "用户列表")
        })
        adminRouters.GET("/user/add", func(c *gin.Context) {
            c.String(http.StatusOK, "增加用户")
        })
        adminRouters.GET("/user/edit", func(c *gin.Context) {
            c.String(http.StatusOK, "修改用户")
        })
        adminRouters.GET("/article", func(c *gin.Context) {
            c.String(http.StatusOK, "新闻列表")
        })
    }
}
