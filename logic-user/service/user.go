package service

import (
	"context"
	"logic-user/dao"
	"logic-user/model"
	"protocols/user"
	"utils"
)

func (i *UserRpcServer) Login(ctx context.Context, req *user.LoginRequest, rsp *user.LoginResponse) error {
	loginUser, err := dao.QueryUser(&model.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		rsp.Code = 201
		return nil
	}
	roles, _ := dao.QueryRoles(loginUser.Id)
	roleIds := make([]int64, 0, len(roles))
	for _, v := range roles {
		roleIds = append(roleIds, v.Id)
	}
	rsp.Token, _ = utils.GenerateToken(loginUser.Id, roleIds)
	rsp.Code = 200
	rsp.User = &user.User{Id: loginUser.Id, Username: loginUser.Username,
		HeadImg: loginUser.HeadImg, Nickname: loginUser.Nickname}
	return nil
}

func (i *UserRpcServer) UserInfo(ctx context.Context, req *user.UserInfoRequest, rsp *user.UserInfoResponse) error {
	find := &model.User{}
	find.Id = req.Uid
	findUser, err := dao.QueryUser(find)
	if err != nil {
		rsp.Code = 201
		return nil
	}
	rsp.Code = 200
	rsp.User = &user.User{Id: findUser.Id, Username: findUser.Username,
		HeadImg: findUser.HeadImg, Nickname: findUser.Nickname}
	return nil
}
