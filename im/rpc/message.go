package rpc

import (
	"context"
	"im/logic"
	"protocols/message"
)

func (i *ImRpcServer) MessageList(ctx context.Context, req *message.MessageListRequest, rsp *message.MessageListResponse) error {
	rsp.Total = 0
	return nil
}

func (i *ImRpcServer) SendMessage(ctx context.Context, req *message.MessageRequest, rsp *message.MessageResponse) error {
	// 1.消息入库
	if err := logic.SaveMessage(req.Message); err != nil {
		rsp.Code = 505
		rsp.Message = "消息保存失败"
		return nil
	}
	rsp.Code = 200
	return nil
}
