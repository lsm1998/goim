package router

import (
	"github.com/gin-gonic/gin"
	http "net/http"
	"strconv"
	"utils"
)

// 鉴权
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		uidStr := c.GetHeader("uid")  // 用户uid
		token := c.GetHeader("token") // 访问令牌
		var errMsg string
		var loginClaims *utils.LoginClaims
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			errMsg = "请求未包含有效用户ID"
			goto FAIT
		}
		loginClaims, err = utils.ValidToken(token, uid)
		if err != nil {
			errMsg = err.Error()
			goto FAIT
		}
		// 验证通过，会继续访问下一个中间件
		c.Set("roles", loginClaims.Role)
		c.Next()
		return
	FAIT:
		// 验证不通过，不再调用后续的函数处理
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"message": errMsg})
	}
}
