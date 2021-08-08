package build

import (
	"common/net/network"
	"common/net/network/TCP"
	"common/net/network/UDP"
	"common/net/network/WS"
	"github.com/panjf2000/gnet"
)

type NetServerBuilder struct {
	network    network.Network
	addr       string
	opts       []gnet.Option
	requestURI string
	initF      func()
	stopF      func()
}

func (n *NetServerBuilder) RequestURI(requestURI string) *NetServerBuilder {
	n.requestURI = requestURI
	return n
}

func (n *NetServerBuilder) Network(network network.Network) *NetServerBuilder {
	n.network = network
	return n
}

func (n *NetServerBuilder) Addr(addr string) *NetServerBuilder {
	n.addr = addr
	return n
}

func (n *NetServerBuilder) Option(opts ...gnet.Option) *NetServerBuilder {
	n.opts = opts
	return n
}

func (n *NetServerBuilder) SetUp(initF func()) *NetServerBuilder {
	n.initF = initF
	return n
}

func (n *NetServerBuilder) SetShutdown(stopF func()) *NetServerBuilder {
	n.stopF = stopF
	return n
}

func (n *NetServerBuilder) Build() network.NetworkServer {
	switch n.network {
	case network.NetworkTCP:
		return TCP.NewTcpNetworkServer(n.addr, n.initF, n.stopF, n.opts...)
	case network.NetworkUDP:
		return UDP.NewUdpNetworkServer(n.addr, n.initF, n.stopF, n.opts...)
	case network.NetworkWebSocket:
		return WS.NewWsNetworkServer(n.addr, n.requestURI, n.initF, n.stopF, n.opts...)
	default:
		panic("network configuration required!")
	}
	return nil
}
