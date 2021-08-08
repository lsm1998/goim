package UDP

import (
	"common/net/network"
	"github.com/panjf2000/gnet"
)

type udpNetworkServer struct {
	network.EventPublisher
	Addr  string
	opts  []gnet.Option
	initF func()
	stopF func()
}

func (t *udpNetworkServer) InitServer() {
	if t.initF != nil {
		t.initF()
	}
}

func (t *udpNetworkServer) StopServer() {
	if t.stopF != nil {
		t.stopF()
	}
}

func (t *udpNetworkServer) StartServer() {
	if err := gnet.Serve(t, "udp://"+t.Addr, t.opts...); err != nil {
		panic(err)
	}
}

func NewUdpNetworkServer(addr string, initF, stopF func(), opts ...gnet.Option) *udpNetworkServer {
	server := new(udpNetworkServer)
	server.Addr = addr
	server.opts = opts
	server.initF = initF
	server.stopF = stopF
	return server
}
