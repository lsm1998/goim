package logic

import (
	"github.com/gin-gonic/gin"
	"web/client"
)

func UserInfo(c *gin.Context) {
	rsp := gin.H{
		"version": 0.1,
		"info":    "hello word!",
	}
	c.JSON(200, rsp)
}

func SendMsg(c *gin.Context) {
	client.ImClient.Auth("")
}
