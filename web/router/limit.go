package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

// 最多每秒允许 limit 个请求
var lmt = rate.NewLimiter(rate.Limit(100), 1)

// 限流
func LimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !lmt.AllowN(time.Now(), 1) {
			c.JSON(200, gin.H{
				"code":    501,
				"message": "服务端繁忙",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
