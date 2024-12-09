package middlewares

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func GlobalMiddlewareOne(c *gin.Context) {
    fmt.Printf("全局-中间件1 --- start\n")
    c.Next()
    fmt.Printf("全局-中间件1 --- end\n")
}

func GlobalMiddlewareTwo(c *gin.Context) {
    fmt.Printf("全局-中间件2 --- start\n")
    c.Next()
    fmt.Printf("全局-中间件2 --- end\n")
}
