package service

import (
	"context"
	proto "protocols/route"
	"route/config"
)

// RouteServer route服务接口
type RouteServer struct {
}

func NewRouteServer() *RouteServer {
	return &RouteServer{}
}

func (r *RouteServer) GetAddr(ctx context.Context, req *proto.GetAddrReq) (*proto.GetAddrReply, error) {
	addr, err := config.EtcdClient.GetAddr(req.ServiceName)
	if err != nil {
		return nil, err
	}
	return &proto.GetAddrReply{
		Code: 200,
		Addr: addr,
	}, nil
}

func (r *RouteServer) GetAddrList(ctx context.Context, req *proto.GetAddrListReq) (*proto.GetAddrListReply, error) {
	list := config.EtcdClient.AddrList(req.ServiceName)
	return &proto.GetAddrListReply{
		Code: 200,
		AddrList: list,
	}, nil
}
