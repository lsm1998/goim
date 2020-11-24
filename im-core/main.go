package main

import (
	"github.com/panjf2000/gnet"
	"im/handler"
	"im/rpc"
	"log"
)

func main() {
	echo := &handler.ImServer{}
	echo.Init()
	defer echo.Stop()
	// rpc服务启动
	go rpc.Init()
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}
