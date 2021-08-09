package strategy

import (
	"comet/handler/route"
	"github.com/panjf2000/gnet"
	"github.com/spf13/cast"
	proto "protocols/message"
)

var (
	handshakeOk = createOkResponse(proto.MessageType_Handshake)
)

type Handshake struct {
}

func (p *Handshake) Handler(msg *proto.Message, c gnet.Conn) error {
	uid := cast.ToInt64(string(msg.Body))
	if uid == 0 {
		return c.AsyncWrite(createFailResponse(proto.MessageType_Handshake, 500, "用户ID错误"))
	}
	if err := route.Join(uid, c); err != nil {
		return c.AsyncWrite(createFailResponse(proto.MessageType_Handshake, 500, err.Error()))
	}
	return c.AsyncWrite(handshakeOk)
}
