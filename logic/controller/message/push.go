package message

import (
	"common"
	"github.com/gin-gonic/gin"
)

type Message struct {
	common.Api
}

func (m *Message) Broadcast(ctx *gin.Context) {

}
