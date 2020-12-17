package rpc

import (
	"context"
	"im/logic"
	"protocols/message"
)

func (i *ImRpcServer) SendMessage(ctx context.Context, req *message.MessageRequest, rsp *message.MessageRequest) error {
	rsp.Type = message.RequestType_Response
	reply := new(message.Reply)
	rsp.Pack = &message.MessageRequest_Response{
		Response: reply,
	}
	if req.Type == message.RequestType_Request {
		switch pack := req.Pack.(type) {
		case *message.MessageRequest_Message:
			// 1.消息入库
			if err := logic.SaveMessage(pack.Message, uint(req.Cmd)); err != nil {
				reply.Code = 505
				reply.Body = []byte("消息保存失败")
				return nil
			}
		case *message.MessageRequest_Response:

		}
	}
	return nil
}
