package client

import (
	"common/etcd"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"logic/config"
	proto "protocols/route"
	"strings"
)

var RouteClient proto.RouteServiceClient

func init() {
	for _, v := range config.C.Servers {
		if v != "im-route" {
			continue
		}
		r := etcd.NewResolver(strings.Join(config.C.Adders, ";"))
		resolver.Register(r)
		//客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
		conn, err := grpc.Dial(r.Scheme()+"://author/"+v, grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
		if err != nil {
			fmt.Println("连接服务器失败：", err)
		}
		RouteClient = proto.NewRouteServiceClient(conn)
	}
}
