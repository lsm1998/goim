package main

import (
	"file_server/config"
	"file_server/route"
	"file_server/server"
)

func main() {
	// 初始化路由
	route.InitRoute()

	// 开始监听
	server.RunHttp(config.HttpAddr())
}
