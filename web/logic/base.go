package logic

import "github.com/gin-gonic/gin"

func Info(c *gin.Context) {
	rsp := gin.H{
		"version": 0.1,
		"info":    "hello word!",
	}
	c.JSON(200, rsp)
}
