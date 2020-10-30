package logic

import (
	"github.com/gin-gonic/gin"
	"utils"
)

func Info(c *gin.Context) {
	rsp := gin.H{
		"version": utils.Version,
		"info":    utils.ProjectName,
	}
	c.JSON(200, rsp)
}
