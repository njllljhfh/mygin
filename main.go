package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "html/template"
    "net/http"
    "os"
    "time"
)

type Article struct {
    Title   string `json:"title"` // 让返回给web的json的key首字母小写
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

// UnixToTime 时间戳转换成日期字符串
func UnixToTime(timestamp int) string {
    fmt.Printf("时间戳为：%v\n", timestamp)
    t := time.Unix(int64(timestamp), 0)
    return t.Format("2006-01-02 15:04:05")
}

func MyPrintln(str1 string, str2 string) string {
    fmt.Println(str1, str2)
    return str1 + "---" + str2
}

func main() {
    // 创建一个默认的路由引擎
    r := gin.Default()

    // 注册自定义的模板函数(注意，注册函数要放在加载模板文件路径之前)
    r.SetFuncMap(template.FuncMap{
        "UnixToTime": UnixToTime,
        "MyPrintln":  MyPrintln,
    })

    // 配置模板文件路径
    // 当模板根目录下有子目录时，每一层子目录都要加上/**
    r.LoadHTMLGlob("templates/**/*")

    // 配置静态web目录
    // 参数1：路由
    // 参数2：映射的目录
    r.Static("/static", "./static")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "default/index.html", gin.H{
            "title": "首页",
            "msg":   "我是msg",
            "score": 89,
            "hobby": []string{"吃饭", "睡觉", "写代码"},
            "newsList": []interface{}{
                &Article{
                    Title:   "新闻-标题111",
                    Desc:    "新闻-描述111",
                    Content: "新闻-详情111",
                },
                &Article{
                    Title:   "新闻-标题222",
                    Desc:    "新闻-描述222",
                    Content: "新闻-详情222",
                },
            },
            "testSlice": []string{},
            "news": &Article{
                Title:   "新闻标题",
                Content: "新闻内容",
            },
            "data": 1733304926, // 时间戳
        })
    })

    r.GET("/news", func(c *gin.Context) {
        news := &Article{
            Title:   "新闻-标题",
            Desc:    "新闻-描述",
            Content: "新闻-详情",
        }
        // 返回html模板
        c.HTML(http.StatusOK, "default/news.html", gin.H{
            "title": "新闻页面",
            "news":  news,
        })
    })

    r.GET("/admin", func(c *gin.Context) {
        c.HTML(http.StatusOK, "admin/index.html", gin.H{})
    })

    r.GET("/admin/news", func(c *gin.Context) {
        // name参数是 html中define定义的名字
        c.HTML(http.StatusOK, "adminNews", gin.H{})
    })
    // - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

    // get请求
    r.GET("/str", func(c *gin.Context) {
        c.String(http.StatusOK, "值，%v", "你好gin")
    })

    r.GET("/json1", func(c *gin.Context) {
        c.JSON(http.StatusOK, map[string]interface{}{
            "success": true,
            "msg":     "json1",
        })
    })

    r.GET("/json2", func(c *gin.Context) {
        // type H map[string]any
        c.JSON(http.StatusOK, gin.H{
            "success": false,
            "msg":     "json2",
        })
    })

    r.GET("/json3", func(c *gin.Context) {
        a := &Article{
            Title:   "标题",
            Desc:    "描述",
            Content: "内容",
        }
        c.JSON(http.StatusOK, a)
    })

    // jsonp请求，可以传入回调函数
    r.GET("/jsonp", func(c *gin.Context) {
        a := &Article{
            Title:   "标题",
            Desc:    "描述",
            Content: "jsonp",
        }
        c.JSONP(http.StatusOK, a)
    })

    // 返回XML数据
    r.GET("/xml", func(c *gin.Context) {
        c.XML(http.StatusOK, gin.H{
            "success": true,
            "msg":     "xml数据",
        })
    })

    // 返回html模板
    r.GET("/goods", func(c *gin.Context) {
        c.HTML(http.StatusOK, "goods.html", gin.H{
            "title": "我是后台数据-商品",
            "price": 20,
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
    // mac下，热启动服务并设置端口，命令行执行: PORT=8086 air
    // windows下， 先执行 $env:PORT="8086"，再执行 air
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // 默认端口
    }
    r.Run(":" + port)
}
