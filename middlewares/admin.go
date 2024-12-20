package middlewares

import (
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "mygin/datamodels"
    "mygin/logconfig"
    "time"
)

var logger *log.Logger

func init() {
    logger = logconfig.InitLogger()
}

func AdminMiddlewareOne(c *gin.Context) {
    // fmt.Println("路由分组admin-中间件1 --- start")
    logger.Infof("路由分组admin-中间件1 --- start")
    // fmt.Println(c.Request.URL)

    // 中间件 与 controller函数 共享数据
    c.Set("userInfo", &datamodels.UserInfo{Name: "老王", Age: 18, Gender: "男"})
    // c.Set("userInfo", "哈哈")  // 用来模拟数据转换失败

    // 当在中间件或handler中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）
    cCp := c.Copy()
    go func() {
        time.Sleep(5 * time.Second)
        // fmt.Printf("协程：Done! in path: %v\n", cCp.Request.URL.Path)
        logger.Infof("协程：Done! in path: %v", cCp.Request.URL.Path)
    }()

    c.Next()
    // fmt.Println("路由分组admin-中间件1 --- end")
    logger.Infof("路由分组admin-中间件1 --- end")
}

func AdminMiddlewareTwo(c *gin.Context) {
    // fmt.Println("路由分组admin-中间件2 --- start")
    logger.Infof("路由分组admin-中间件2 --- start")
    c.Next()
    // fmt.Println("路由分组admin-中间件2 --- end")
    logger.Infof("路由分组admin-中间件2 --- start")
}
