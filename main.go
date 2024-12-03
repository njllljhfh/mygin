package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func main() {
    // 创建一个默认的路由引擎
    r := gin.Default()

    // get请求
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "值，%v", "你好gin")
    })

    r.GET("/news", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "newsName":    "吃大瓜...",
            "newsContent": "我是新闻页面 111",
        })
    })

    r.POST("/add", func(c *gin.Context) {
        c.String(http.StatusOK, "我是post返回的数据-%d", 666)
    })

    r.DELETE("/delete", func(c *gin.Context) {
        c.String(http.StatusOK, "我是delete返回的数据-%d", 777)
    })

    // 启动HTTP服务 默认在 0.0.0.0:8080 上启动服务
    // r.Run(":8686") // 设置端口为 8686
    // - - -
    // 热启动服务并设置端口，命令行执行: PORT=8086 air
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // 默认端口
    }
    r.Run(":" + port)
}
