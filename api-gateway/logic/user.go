package logic

import (
	"api-gateway/client"
	"github.com/gin-gonic/gin"
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
