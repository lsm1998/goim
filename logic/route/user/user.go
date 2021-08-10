package user

import (
	"github.com/gin-gonic/gin"
	"logic/controller/user"
)

func InitUserRoute(r *gin.RouterGroup) {
	group := r.Group("/user")
	{
		u := user.User{}
		group.POST("/login", u.Login)
	}
}
