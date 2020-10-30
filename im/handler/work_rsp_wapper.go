package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/panjf2000/gnet"
	"protocols/message"
)

func replyMessage(c gnet.Conn, reply *message.Reply) error {
	var rsp message.MessageRequest
	rsp.Type = message.RequestType_Response
	rsp.Pack = &message.MessageRequest_Response{Response: reply}
	if data, err := proto.Marshal(&rsp); err == nil {
		return c.AsyncWrite(data)
	}
	return nil
}
