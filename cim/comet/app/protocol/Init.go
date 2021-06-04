package protocol

import (
	"utils/chain"
)

// Register 注册协议转换器
func Register() {
	registerProtocol()
}

// registerProtocol 注册一次协议转换器
func registerProtocol() {
	var transformLine chain.TransformLine
	transformLine = newUDPTransformLine()
	if err := chain.ProtocolRegister(transformLine); err != nil {
		panic(err)
	}
}
