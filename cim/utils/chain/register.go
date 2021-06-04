package chain

// ProtocolRegister 协议转换器注册
func ProtocolRegister(transform TransformLine) error {
	protocolSink = append(protocolSink, transform)
	return nil
}

// EventRegister 事件处理器注册
// note 排序ID相同时，调用顺序不可预测
func EventRegister(sortID int, f func() EventHandler) error {
	create = append(create, createStruct{
		f, sortID,
	})
	return nil
}

// SinkRegister 输出注册（所有输出不允许再次修改line)
func SinkRegister(sink Sink) error {
	sinks = append(sinks, sink)
	return nil
}

// NetWorkRegister 输出network register
func NetWorkRegister(create func() NetWork) {
	netWorks = create
}
