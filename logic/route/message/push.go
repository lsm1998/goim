package message

import (
	"github.com/gin-gonic/gin"
	"logic/controller/message"
)

func InitMessageRoute(r *gin.RouterGroup) {
	r.Group("/push")
	{
		m := &message.Message{}
		r.GET("/broadcast", m.Broadcast)
	}
}
