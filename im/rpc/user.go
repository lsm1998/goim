package rpc

import (
	"context"
	"im/dao"
	"im/route"
	"protocols/user"
	"utils"
)

func (i *ImRpcServer) UserList(ctx context.Context, req *user.UserListRequest, rsp *user.UserListResponse) error {
	// 在线列表
	list := route.KeyList()
	pageInfo := &utils.PageInfo{
		Page: req.Page,
		Size: req.Size,
	}
	userList, err := dao.QueryUserList(list, pageInfo)
	if err != nil {
		return err
	}
	rsp.Total = userList.Total
	rsp.List = make([]*user.User, 0, req.Size)
	for _, v := range userList.List {
		rsp.List = append(rsp.List, &user.User{
			Id:       v.Id,
			Nickname: v.Nickname,
			Username: v.Username,
			HeadImg:  v.HeadImg,
		})
	}
	return nil
}
