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
	msg := &proto.MessageRequest{}
	err := protobuf.Unmarshal(data, msg)
	if err != nil {
		log.Error(err)
		return
	}
	switch msg.Pack.(type) {
	case *proto.MessageRequest_Response:
		if msg.Type != proto.RequestType_Response {
			log.Error("type need Response")
			return
		}
		processorResponse(msg.Cmd, msg.Pack.(*proto.MessageRequest_Response), c)
	case *proto.MessageRequest_Message:
		if msg.Type != proto.RequestType_Request {
			log.Error("type need Request")
			return
		}
		processorMessage(msg.Cmd, msg.Pack.(*proto.MessageRequest_Message), c)
	default:
		log.Error("msg.Pack type not support")
		return
	}
}



func (*ReactHandler) EventType() network.NetWorkEventType {
	return network.NetWorkEventReact
}
