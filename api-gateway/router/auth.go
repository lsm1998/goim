package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	http "net/http"
	"strconv"
	"utils"
)

// JWT鉴权
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		uidStr := c.GetHeader("uid")  // 用户uid
		token := c.GetHeader("token") // 访问令牌
		var errMsg string
		var loginClaims *utils.LoginClaims
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			errMsg = "请求未包含有效用户ID"
			goto FAIL
		}
		loginClaims, err = utils.ValidToken(token, uid)
		if err != nil {
			errMsg = err.Error()
			goto FAIL
		}
		// 验证通过，会继续访问下一个中间件
		c.Set("roles", loginClaims.Role)
		c.Set("uid", loginClaims.Uid)
		c.Next()
		return
	FAIL:
		// 验证不通过，不再调用后续的函数处理
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": errMsg})
	}
}

// 路径验证
func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 访问的页面
		path := c.Request.URL.Path
		fmt.Println(path)
		c.Next()
		// FAIT:
		// 验证不通过，不再调用后续的函数处理
		//c.Abort()
		//c.JSON(http.StatusUnauthorized, gin.H{"message": errMsg})
	}
}
