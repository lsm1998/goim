package server

import (
	"utils/chain"
)

func Register() {
	registerOnLineServer()
}

// registerOnLineServer 注册一次网络服务
func registerOnLineServer() {
	chain.NetWorkRegister(func() chain.NetWork {
		return newIMServer()
	})
}
