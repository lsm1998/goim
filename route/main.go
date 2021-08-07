package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	proto "protocols/route"
	"route/config"
	"route/service"
)

func main() {
	host := fmt.Sprintf("127.0.0.1:%d", config.C.Rpc.Port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// 创建grpc句柄
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	// 将greetServer结构体注册到grpc服务中
	proto.RegisterRouteServiceServer(srv, service.NewRouteServer())

	// 将服务地址注册到etcd中
	err = config.EtcdClient.Register(config.C.Rpc.Server, host)
	if err != nil {
		panic(err)
	}

	// 监听服务
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
