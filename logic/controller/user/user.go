package user

import (
	"common"
	"github.com/gin-gonic/gin"
	"logic/model/user"
	"logic/service"
)

type User struct {
	common.Api
	userService service.UserService
}

func (u *User) Login(ctx *gin.Context) {
	user := user.User{}
	if err := ctx.BindJSON(&user); err != nil {
		common.Error(ctx, 400, "参数解析错误")
		return
	}
	addr, err := u.userService.Login(&user)
	if err != nil {
		common.Error(ctx, 500, err.Error())
		return
	}
	common.OK(ctx, addr, "ok")
}
