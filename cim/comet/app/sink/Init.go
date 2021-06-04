package sink

import (
	"utils/chain"
)

func Register() {
	registerOnLineServer()
}

// registerOnLineServer 注册一次网络服务
func registerOnLineServer() {
	if err := chain.SinkRegister(newFileDataSink()); err != nil {
		panic(err)
	}
}
