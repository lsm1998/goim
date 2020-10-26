package rpc

import (
	"context"
	"protocols/message"
)

func (i *ImRpcServer) MessageList(ctx context.Context, req *message.MessageListRequest, rsp *message.MessageListResponse) error {
	rsp.Total = 0

	return nil
}

func (i *ImRpcServer) SendMessage(ctx context.Context, req *message.MessageRequest, rsp *message.MessageResponse) error {
	return nil
}
