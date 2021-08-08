package TCP

import (
	"common/net/network"
	"github.com/panjf2000/gnet"
)

type tcpNetworkServer struct {
	network.EventPublisher
	Addr         string
	opts         []gnet.Option
	initF, stopF func()
}

func (t *tcpNetworkServer) InitServer() {
	if t.initF != nil {
		t.initF()
	}
}

func (t *tcpNetworkServer) StopServer() {
	if t.stopF != nil {
		t.stopF()
	}
}

func (t *tcpNetworkServer) StartServer() {
	if err := gnet.Serve(t, "tcp://"+t.Addr, t.opts...); err != nil {
		panic(err)
	}
}

func NewTcpNetworkServer(addr string, initF, stopF func(), opts ...gnet.Option) *tcpNetworkServer {
	server := new(tcpNetworkServer)
	server.Addr = addr
	server.opts = opts
	server.initF, server.stopF = initF, stopF
	return server
}
