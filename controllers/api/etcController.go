package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EtcController struct{}

func (con *EtcController) UserList(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api接口-userlist")
}

func (con *EtcController) Plist(c *gin.Context) {
	c.String(http.StatusOK, "我是一个api接口-plist")
}
