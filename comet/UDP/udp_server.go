package UDP

import (
	"comet"
	"github.com/panjf2000/gnet"
)

type udpNetworkServer struct {
	comet.EventPublisher
	Addr string
	opts []gnet.Option
}

func (*udpNetworkServer) InitServer() {

}

func (*udpNetworkServer) StopServer() {

}

func (t *udpNetworkServer) StartServer() {
	_ = gnet.Serve(t, "udp://"+t.Addr, t.opts...)
}

func NewUdpNetworkServer(addr string, opts ...gnet.Option) *udpNetworkServer {
	server := new(udpNetworkServer)
	server.Addr = addr
	server.opts = opts
	return server
}
