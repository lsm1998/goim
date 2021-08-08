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

// AddrList 返回当前服务下可用的地址列表
func (c *Comet) AddrList(ctx *gin.Context) {
	server := ctx.Query("server")
	if server == "" {
		common.Error(ctx, 100, "缺少必填参数")
		return
	}
	req := &proto.GetAddrListReq{ServiceName: server}
	data, err := client.RouteClient.GetAddrList(ctx, req)
	if err != nil {
		common.Error(ctx, 100, err.Error())
		return
	}
	common.OK(ctx, data.AddrList, "ok")
}

// Addr 返回当前服务下一个可用的服务，基于最小连接
func (c *Comet) Addr(ctx *gin.Context) {
	server := ctx.Query("server")
	if server == "" {
		common.Error(ctx, 100, "缺少必填参数")
		return
	}
	req := &proto.GetAddrReq{ServiceName: server}
	data, err := client.RouteClient.GetAddr(ctx, req)
	if err != nil {
		common.Error(ctx, 100, err.Error())
		return
	}
	common.OK(ctx, data, "ok")
}
