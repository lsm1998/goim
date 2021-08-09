package strategy

import (
	"comet/handler/route"
	"github.com/panjf2000/gnet"
	"github.com/prometheus/common/log"
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
	route.Foreach(func(c gnet.Conn) {
		if err := c.AsyncWrite(data); err != nil {
			log.Error(err)
		}
	})
	return nil
}
