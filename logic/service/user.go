package service

import (
	"context"
	"logic/client"
	"logic/model/user"
	proto "protocols/route"
)

const ImServerName = "im-comet"

type UserService struct {
}

func NewUserService() *UserService {
	return nil
}

func (u *UserService) Login(user *user.User) (string, error) {
	// 返回一个可用的服务器地址
	reply, err := client.RouteClient.GetAddr(context.TODO(), &proto.GetAddrReq{
		ServiceName: ImServerName,
	})
	if err != nil {
		return "", err
	}
	return reply.Addr, nil
}
