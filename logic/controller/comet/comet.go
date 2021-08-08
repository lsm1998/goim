package comet

import (
	"common"
	"github.com/gin-gonic/gin"
	"logic/client"
	proto "protocols/route"
)

type Comet struct {
	common.Api
}

func (c *Comet) AddrList(ctx *gin.Context) {
	req := &proto.GetAddrListReq{}
	list, err := client.RouteClient.GetAddrList(ctx, req)
	if err != nil {
		c.Error(100, err, err.Error())
		return
	}
	c.OK(list, "ok")
}

func (c *Comet) Addr(ctx *gin.Context) {
	req := &proto.GetAddrReq{}
	data, err := client.RouteClient.GetAddr(ctx, req)
	if err != nil {
		c.Error(100, err, err.Error())
		return
	}
	c.OK(data, "ok")
}
