package client

import (
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	ImClient   client.XClient
	UserClient client.XClient
)

func Init() {
	{
		dis := client.NewConsulDiscovery("IM", "IM", []string{"47.103.211.234:8500"}, nil)
		option := client.DefaultOption
		option.SerializeType = protocol.ProtoBuffer
		ImClient = client.NewXClient("IM", client.Failtry, client.RoundRobin, dis, option)
	}
	{
		dis := client.NewConsulDiscovery("USER", "USER", []string{"47.103.211.234:8500"}, nil)
		option := client.DefaultOption
		option.SerializeType = protocol.ProtoBuffer
		UserClient = client.NewXClient("USER", client.Failtry, client.RoundRobin, dis, option)
	}
}
