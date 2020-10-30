package client

import (
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	ImClient client.XClient
)

func Init() {
	dis := client.NewConsulDiscovery("IM", "IM", []string{"47.103.211.234:8500"}, nil)
	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer
	ImClient = client.NewXClient("IM", client.Failtry, client.RoundRobin, dis, option)
}
