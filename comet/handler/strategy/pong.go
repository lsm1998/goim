package strategy

import (
	"comet/handler/route"
	"github.com/panjf2000/gnet"
	proto "protocols/message"
)

var (
	pongOk = createOkResponse(proto.MessageType_Pong)
)

type Pong struct {
}

func (p *Pong) Handler(msg *proto.Message, c gnet.Conn) error {
	if err := route.Pong(msg.FormId); err != nil {
		return c.AsyncWrite(createFailResponse(proto.MessageType_Pong, 500, err.Error()))
	}
	return c.AsyncWrite(pongOk)
}
