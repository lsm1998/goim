package handler

import (
	"common/net/network"
	"github.com/panjf2000/gnet"
)

type CloseHandler struct {
}

func (*CloseHandler) Handler(data *[]byte, c gnet.Conn) {
	connMaps.LeaveByConn(c)
}

func (*CloseHandler) EventType() network.NetWorkEventType {
	return network.NetWorkEventClose
}