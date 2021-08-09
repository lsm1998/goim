package handler

import (
	"comet/handler/route"
	"common/net/network"
	"github.com/panjf2000/gnet"
)

type CloseHandler struct {
}

func (*CloseHandler) Handler(data []byte, c gnet.Conn) {
	route.connMaps.LeaveByConn(c)
}

func (*CloseHandler) EventType() network.NetWorkEventType {
	return network.NetWorkEventClose
}
