package itying

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "mygin/models"
    "net/http"
    "time"
)

type IndexController struct{}

type Article struct {
    Title   string `json:"title"` // 让返回给web的json的key首字母小写
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

// Index 前台首页
func (con *IndexController) Index(c *gin.Context) {
    // 获取当前时间
    now := time.Now()
    // 秒级时间戳
    timestamp := now.Unix()
    fmt.Println("秒级时间戳:", timestamp)
    // 调用自定义模块中的函数
    fmt.Println("秒级日期:", models.UnixToTime(int(timestamp)))

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
        "data": int(timestamp), // 时间戳, html模板中调用了 UnixToTime 方法
    })
}
