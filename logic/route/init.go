package route

import (
	"github.com/gin-gonic/gin"
	"logic/route/comet"
	"logic/route/message"
	"logic/route/monitor"
	"logic/route/user"
)

func InitRoute(r *gin.RouterGroup) {
	monitor.InitMonitorRoute(r)
	message.InitMessageRoute(r)
	comet.InitCometRoute(r)
	user.InitUserRoute(r)
}
