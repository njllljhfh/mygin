package admin

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController // 结构体继承
}

func (con *UserController) Index(c *gin.Context) {
	//c.String(http.StatusOK, "用户列表")
	con.success(c, "用户列表---Index")
}

func (con *UserController) Add(c *gin.Context) {
	//c.String(http.StatusOK, "增加用户")
	con.success(c, "用户列表---Add")
}

func (con *UserController) Edit(c *gin.Context) {
	//c.String(http.StatusOK, "修改用户")
	con.error(c, "用户列表---Edit")
}
