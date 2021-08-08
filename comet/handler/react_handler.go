package handler

import (
	"common/net/network"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/panjf2000/gnet"
	"github.com/prometheus/common/log"
	proto "protocols/message"
)

type ReactHandler struct {
}

func (*ReactHandler) Handler(data []byte, c gnet.Conn) {
	meg := &proto.Message{}
	err := protobuf.Unmarshal(data, meg)
	if err != nil {
		log.Error(err)
		return
	}
}

func (*ReactHandler) EventType() network.NetWorkEventType {
	return network.NetWorkEventReact
}
