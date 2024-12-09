package admin

import (
    "fmt"
    "github.com/gin-gonic/gin"
    dataModel "mygin/datamodals"
)

type EmptyJson struct{}

type Profile struct {
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=1"`
}

type UserAddRequest struct {
    Name    string  `json:"name" binding:"required"`
    Profile Profile `json:"profile" binding:"required"`
}

type UserController struct {
    BaseController // 结构体继承
}

func (con *UserController) User(c *gin.Context) {
    // 获取中间件中添加的数据
    // 注意：user 的类型是 空接口
    user, exists := c.Get("userInfo")
    if !exists {
        msg := "userInfo 不存在"
        con.error(c, msg, EmptyJson{})
        return
    }
    // 类型断言，类型类型装换
    v, ok := user.(*dataModel.UserInfo)
    if !ok {
        msg := "userInfo 转换失败"
        con.error(c, msg, EmptyJson{})
        return
    }
    con.success(c, "用户列表---Index", v)
}

// UserAdd 获取POST传递的嵌套json，绑定json数据到自定义的结构体
func (con *UserController) UserAdd(c *gin.Context) {
    req := UserAddRequest{}
    if err := c.ShouldBind(&req); err != nil {
        errMsg := fmt.Sprintf("用户列表---Add: %v", err)
        fmt.Println(errMsg)
        con.error(c, errMsg, EmptyJson{})
        return
    }

    fmt.Printf("%#v\n", req)
    con.success(c, "用户列表---Add", gin.H{
        "user": req,
    })
}

func (con *UserController) UserEdit(c *gin.Context) {
    con.error(c, "用户列表---Edit", EmptyJson{})
}
