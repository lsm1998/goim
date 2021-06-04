package TCP

import (
	"github.com/panjf2000/gnet"
	"utils/network"
)

type tcpNetworkServer struct {
	network.EventPublisher
	Addr string
	opts []gnet.Option
}

func (*tcpNetworkServer) InitServer() {

}

func (*tcpNetworkServer) StopServer() {

}

func (t *tcpNetworkServer) StartServer() {
	_ = gnet.Serve(t, "tcp://"+t.Addr, t.opts...)
}

func NewTcpNetworkServer(addr string, opts ...gnet.Option) *tcpNetworkServer {
	server := new(tcpNetworkServer)
	server.Addr = addr
	server.opts = opts
	return server
}
