package strategy

import (
	"github.com/panjf2000/gnet"
	proto "protocols/message"
)

type Broadcast struct {
}

func (p *Broadcast) Handler(msg *proto.Message, c gnet.Conn) error {
	data := createResponse(proto.MessageType_SystemBroadcast, &proto.Reply{
		Code:   200,
		Body:   msg.Body,
		FormId: msg.FormId,
		MsgId:  0,
	})
	return c.AsyncWrite(data)
}
