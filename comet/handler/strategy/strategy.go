package strategy

import (
	"github.com/panjf2000/gnet"
	proto "protocols/message"
)

type MsgHandlerStrategy interface {
	Handler(message *proto.Message, c gnet.Conn) error
}

type MsgHandler struct {
	context  *MsgHandlerContext
	strategy MsgHandlerStrategy
}

type MsgHandlerContext struct {
	message *proto.Message
	c       gnet.Conn
}

func NewMsgHandler(message *proto.Message, c gnet.Conn, strategy MsgHandlerStrategy) *MsgHandler {
	return &MsgHandler{
		context: &MsgHandlerContext{
			message: message,
			c:       c,
		},
		strategy: strategy,
	}
}

func (p *MsgHandler) Handler() error {
	return p.strategy.Handler(p.context.message, p.context.c)
}
