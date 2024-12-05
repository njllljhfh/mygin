package main

import (
    "encoding/xml"
    "fmt"
    "github.com/gin-gonic/gin"
    "html/template"
    "mygin/routers"
    "net/http"
    "os"
    "time"
)

type Article struct {
    Title   string `json:"title"` // 让返回给web的json的key首字母小写
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

type UserInfo struct {
    Username string `json:"username" form:"username"` // form:"username" 对应web传过来的key
    Password string `json:"password" form:"password"`
}

// Article2 解析xml数据到结构体
type Article2 struct {
    Title   string `json:"title" xml:"title"`
    Content string `json:"content" xml:"content"`
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

    //
    routers.AdminRouterInit(r)
    routers.ApiRouterInit(r)
    routers.DefaultRouterInit(r)
    // - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

    /*
       获取请求中传递的参数
    */

    // get请求传值
    r.GET("/getTest", func(c *gin.Context) {
        username := c.Query("username")
        age := c.Query("age")
        // 如果没有page，设置默认值为 "1"
        page := c.DefaultQuery("page", "1")
        c.JSON(http.StatusOK, gin.H{
            "username": username,
            "age":      age,
            "page":     page,
        })
    })

    r.GET("/user", func(c *gin.Context) {
        c.HTML(http.StatusOK, "default/user.html", gin.H{})
    })
    // post请求传值
    r.POST("/doAddUser", func(c *gin.Context) {
        // 获取表单post请求传递的数据
        username := c.PostForm("username")
        password := c.PostForm("password")
        c.JSON(http.StatusOK, gin.H{
            "username": username,
            "password": password,
        })
    })

    // 把web传递过来的数据绑定到结构体上
    // http://localhost:8086/getUser?username=zhangsan&password=123
    r.GET("/getUser", func(c *gin.Context) {
        user := &UserInfo{}
        // 绑定web传来的数据到结构体
        if err := c.ShouldBind(user); err == nil {
            fmt.Printf("用户结构体信息：%#v\n", user)
            c.JSON(http.StatusOK, user)
        } else {
            c.JSON(http.StatusOK, gin.H{
                "err": err,
            })
        }
    })

    r.POST("/doAddUser2", func(c *gin.Context) {
        user := &UserInfo{}
        // 绑定web传来的数据到结构体
        if err := c.ShouldBind(user); err == nil {
            fmt.Printf("用户结构体信息：%#v\n", user)
            c.JSON(http.StatusOK, user)
        } else {
            c.JSON(http.StatusOK, gin.H{
                "err": err,
            })
        }
    })

    // 获取 Post xml 数据
    // 在对接 第三方提供的支付相关功能时，可能返回的是xml数据
    r.POST("/xml", func(c *gin.Context) {
        article := &Article2{}
        xmlSliceData, _ := c.GetRawData() //  从 c.Request.Body 读取请求数据
        // fmt.Println("xmlSliceData", xmlSliceData)  // []byte
        // 将xml数据绑定到结构体
        if err := xml.Unmarshal(xmlSliceData, article); err == nil {
            c.JSON(http.StatusOK, article)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{
                "err": err,
            })
        }
    })
    // - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

    /*
       动态路由，从路由中获取参数
    */
    // list/123
    // cid=123
    r.GET("/list/:cid", func(c *gin.Context) {
        cid := c.Param("cid")
        fmt.Printf("cid=%s\n", cid)
        c.String(http.StatusOK, "cid: %s", cid)
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
    // - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

    /*
       路由分组, 见 /routers 目录
    */

    // - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

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
