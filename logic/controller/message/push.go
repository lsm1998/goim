package message

import (
	"common"
	"github.com/gin-gonic/gin"
	"logic/service"
)

type Message struct {
	common.Api
	push *service.PushService
}

func NewMessage() *Message {
	return &Message{
		push: service.NewPushService(),
	}
}

func (m *Message) Broadcast(ctx *gin.Context) {
	content := ctx.Query("msg")
	if content == "" {
		common.Error(ctx, 400, "不可以发送空内容")
		return
	}
	// uid := 解析token
	if err := m.push.Broadcast(1, content); err != nil {
		common.Error(ctx, 500, err.Error())
		return
	}
	common.OK(ctx, nil, "ok")
}
