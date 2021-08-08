package chain

import (
	"github.com/panjf2000/gnet"
)

type StopEvent interface {
	// Stop 退出通知（必需是阻塞实现)
	Stop()
}

// EventHandler 事件处理接口
type EventHandler interface {
	StopEvent
	//Filter 数据包过滤,
	// return true:数据正常，false数据包过滤
	Filter(line EventLine) bool

	//Handler 事件接受判断
	Handler(line EventLine) bool
}

// Transforms 协议转换接口
type Transforms interface {
	Line(data *[]byte, conn gnet.Conn) error
	StopEvent
}

// TransformLine 单行转换(需要阻塞实现)
type TransformLine interface {
	Line(data *[]byte, conn gnet.Conn) (line EventLine, err error)
}

// Sink 数据结点
type Sink interface {
	//PushLine 事件推送
	PushLine(line EventLine) error

	// Setup 初始化处理
	Setup() error

	StopEvent
}

// NetWork 网络顶层接口
type NetWork interface {
	// Setup 初始化处理
	Setup(f func(data *[]byte, con gnet.Conn) error) error

	Run()

	StopEvent
}

// EventLine 事件行
type EventLine interface {
}
