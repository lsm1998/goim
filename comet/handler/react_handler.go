package handler

import (
	"common/net/network"
	"fmt"
	"github.com/panjf2000/gnet"
)

type ReactHandler struct {
}

func (*ReactHandler) Handler(data *[]byte, c gnet.Conn) {
	fmt.Println(string(*data))
	_ = c.AsyncWrite([]byte("echo:" + string(*data)))
}

func (*ReactHandler) EventType() network.NetWorkEventType {
	return network.NetWorkEventReact
}
