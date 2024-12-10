package admin

import (
    "fmt"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "mygin/datamodels"
    "mygin/logconfig"
    "net/http"
    "path"
)

var logger *log.Logger

func init() {
    logger = logconfig.InitLogger()
}

type UserController struct {
    BaseController // 结构体继承
}

type EmptyJson struct{}

type Profile struct {
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=1"`
}

type UserAddRequest struct {
    Name    string  `json:"name" binding:"required"`
    Profile Profile `json:"profile" binding:"required"`
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
    v, ok := user.(*datamodels.UserInfo)
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

func (con *UserController) UserAdd2(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/userAdd2.html", gin.H{})
}

// Upload 单文件上传文件
func (con *UserController) Upload(c *gin.Context) {
    username := c.PostForm("username")
    // fmt.Printf("用户名：%v\n", username)
    logger.Infof("用户名：%v", username)

    // 单文件
    file, err := c.FormFile("face")
    if err != nil {
        con.error(c, fmt.Sprintf("%v", err), EmptyJson{})
        return
    }

    logger.WithFields(log.Fields{
        "filename": file.Filename,
    }).Infof("文件名：%v", file.Filename)

    // 上传文件至指定的完整文件路径
    dst := path.Join("./uploadFiles", file.Filename)
    _ = c.SaveUploadedFile(file, dst)
    con.success(c, "上传文件", gin.H{
        "username": username,
        "dst":      dst,
    })

    // fmt.Println("999-这行代码在上面返回成功后，依然会执行") // 这行代码在上面返回成功后，依然会执行
    logger.Infof("999-这行代码在上面返回成功后，依然会执行")
}

func (con *UserController) UserAdd3(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/userAdd3.html", gin.H{})
}

// MultiUpload 多文件上传文件
func (con *UserController) MultiUpload(c *gin.Context) {
    username := c.PostForm("username")
    fmt.Printf("用户名：%v\n", username)

    // 获取
    form, err := c.MultipartForm()
    if err != nil {
        con.error(c, fmt.Sprintf("%v", err), EmptyJson{})
        return
    }

    files, ok := form.File["face[]"]
    if !ok {
        con.error(c, fmt.Sprintf("%v", err), EmptyJson{})
        return
    }

    slice := make([]map[string]string, 0)
    for _, file := range files {
        dst := path.Join("./uploadFiles", file.Filename)
        _ = c.SaveUploadedFile(file, dst)
        slice = append(slice, map[string]string{
            "fileName": file.Filename,
            "filePath": dst,
        })
    }

    con.success(c, "上传文件", gin.H{
        "username":  username,
        "filesInfo": slice,
    })
}
