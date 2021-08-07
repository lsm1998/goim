package service

import (
	"context"
	proto "protocols/route"
)

// RouteServer route服务接口
type RouteServer struct {
}

func NewRouteServer() *RouteServer {
	return &RouteServer{}
}

func (r *RouteServer) GetAddr(ctx context.Context, req *proto.GetAddrReq) (*proto.GetAddrReply, error) {
	return &proto.GetAddrReply{
		Code: 200,
	}, nil
}

func (r *RouteServer) GetAddrList(ctx context.Context, req *proto.GetAddrListReq) (*proto.GetAddrListReply, error) {
	return &proto.GetAddrListReply{
		Code: 200,
	}, nil
}
