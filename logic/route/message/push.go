package message

import (
	"github.com/gin-gonic/gin"
	"logic/controller/message"
)

func InitMessageRoute(r *gin.RouterGroup) {
	group := r.Group("/push")
	{
		m := message.NewMessage()
		group.GET("/broadcast", m.Broadcast)
	}
}
