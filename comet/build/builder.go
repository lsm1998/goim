package build

import (
	"comet"
	"comet/TCP"
	"comet/UDP"
	"comet/WS"
	"github.com/panjf2000/gnet"
)

type NetServerBuilder struct {
	network    comet.Network
	addr       string
	opts       []gnet.Option
	requestURI string
}

func (n *NetServerBuilder) RequestURI(requestURI string) *NetServerBuilder {
	n.requestURI = requestURI
	return n
}

func (n *NetServerBuilder) Network(network comet.Network) *NetServerBuilder {
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

func (n *NetServerBuilder) Build() comet.NetworkServer {
	switch n.network {
	case comet.NetworkTCP:
		return TCP.NewTcpNetworkServer(n.addr, n.opts...)
	case comet.NetworkUDP:
		return UDP.NewUdpNetworkServer(n.addr, n.opts...)
	case comet.NetworkWebSocket:
		return WS.NewWsNetworkServer(n.addr, n.requestURI, n.opts...)
	default:
		panic("network configuration required!")
	}
	return nil
}
