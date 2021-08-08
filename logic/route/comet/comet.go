package comet

import (
	"github.com/gin-gonic/gin"
	"logic/controller/comet"
)

func InitCometRoute(r *gin.RouterGroup) {
	r.Group("/comet")
	{
		c := &comet.Comet{}
		r.GET("/addrList", c.AddrList)
		r.GET("/addr", c.Addr)
	}
}
