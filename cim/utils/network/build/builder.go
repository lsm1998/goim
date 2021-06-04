package build

import (
	"github.com/panjf2000/gnet"
	"utils/network"
	TCP2 "utils/network/TCP"
	UDP2 "utils/network/UDP"
	WS2 "utils/network/WS"
)

type NetServerBuilder struct {
	network    network.Network
	addr       string
	opts       []gnet.Option
	requestURI string
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

func (n *NetServerBuilder) Build() network.NetworkServer {
	switch n.network {
	case network.NetworkTCP:
		return TCP2.NewTcpNetworkServer(n.addr, n.opts...)
	case network.NetworkUDP:
		return UDP2.NewUdpNetworkServer(n.addr, n.opts...)
	case network.NetworkWebSocket:
		return WS2.NewWsNetworkServer(n.addr, n.requestURI, n.opts...)
	default:
		panic("network configuration required!")
	}
	return nil
}
