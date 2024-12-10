package middlewares

import (
    "github.com/gin-gonic/gin"
    "time"
)

/*
在api接口中配置的中间件
*/

func InitMiddlewareOne(c *gin.Context) {
    start := time.Now().UnixNano()
    // fmt.Println("接口级-中间件1 --- start")
    logger.Infof("接口级-中间件1 --- start")

    // 调用该请求的剩余处理程序
    c.Next()

    // 终止调用该请求的剩余处理程序
    // c.Abort()

    // c.Abort() 不会停止当前处理函数，下面的逻辑还会继续执行
    // fmt.Println("接口级-中间件1 --- end")
    logger.Infof("接口级-中间件1 --- end")
    end := time.Now().UnixNano()
    // fmt.Printf("耗时：%v 纳秒\n", end-start)
    logger.Infof("耗时：%v 纳秒\n", end-start)
}

func InitMiddlewareTwo(c *gin.Context) {
    // fmt.Println("接口级-中间件2 --- start")
    logger.Infof("接口级-中间件2 --- start")
    c.Next()
    // fmt.Println("接口级-中间件2 --- end")
    logger.Infof("接口级-中间件2 --- end")
}
