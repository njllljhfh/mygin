package middlewares

import (
    "github.com/gin-gonic/gin"
)

func GlobalMiddlewareOne(c *gin.Context) {
    // fmt.Printf("全局-中间件1 --- start\n")
    logger.Infof("全局-中间件1 --- start")
    c.Next()
    // fmt.Printf("全局-中间件1 --- end\n")
    logger.Infof("全局-中间件1 --- end")
}

func GlobalMiddlewareTwo(c *gin.Context) {
    // fmt.Printf("全局-中间件2 --- start\n")
    logger.Infof("全局-中间件2 --- start")
    c.Next()
    // fmt.Printf("全局-中间件2 --- end\n")
    logger.Infof("全局-中间件2 --- end")
}
