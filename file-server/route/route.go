package route

import (
	"file_server/server"
	"file_server/uphand"
)

// 注册访问路由
func InitRoute() {
	// 路由处理绑定
	server.Handle("/", uphand.Controller{})

	// 获取图片信息
	server.HandleFunc("/info", uphand.Info)

	// 测试上传
	server.HandleFunc("/test", uphand.Test)

	// 获取状态码
	server.HandleFunc("/statuscode", uphand.StatusCode)
}
