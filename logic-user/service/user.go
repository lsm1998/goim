package service

import (
	"context"
	"logic-user/dao"
	"logic-user/model"
	"protocols/user"
	"utils"
)

// Login 登录
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

// UserInfo 用户信息
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

// FriendsList 好友列表
func (i *UserRpcServer) FriendsList(ctx context.Context, req *user.FriendsListRequest, rsp *user.FriendsListResponse) error {
	list, err := dao.QueryFriendsList(&model.Friends{UserId: req.UserId})
	if err != nil {
		rsp.Code = 202
		return nil
	}
	rsp.List = make([]*user.Friends, 0, len(list))
	for _, v := range list {
		temp := &user.Friends{
			Id:      v.FriendsId,
			GroupId: v.GroupId,
		}
		temp.List = make([]*user.User, 0, len(v.List))
		for _, k := range v.List {
			temp.List = append(temp.List, &user.User{
				Id:       k.Id,
				Nickname: k.Nickname,
				HeadImg:  k.HeadImg,
				Username: k.Username,
			})
		}
		rsp.List = append(rsp.List, temp)
	}
	return nil
}
