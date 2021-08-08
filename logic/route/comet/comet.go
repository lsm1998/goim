package comet

import (
	"github.com/gin-gonic/gin"
	"logic/controller/comet"
)

func InitCometRoute(r *gin.RouterGroup) {
	group := r.Group("/comet")
	{
		c := &comet.Comet{}
		group.GET("/addrList", c.AddrList)
		group.GET("/addr", c.Addr)
	}
}
