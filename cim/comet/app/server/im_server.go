package server

import (
	"comet/common/config"
	"fmt"
	"github.com/panjf2000/gnet"
	_ "net/http/pprof"
	"utils/network"
	"utils/network/build"
)

type IMServer struct {
	netServer network.NetworkServer
	handlerF  func(data *[]byte, con gnet.Conn) error
}

func newIMServer() *IMServer {
	instance := new(IMServer)
	builder := &build.NetServerBuilder{}
	cfg := config.NetworkConfig
	switch cfg.Protocol {
	case "TCP":
		builder = builder.Network(network.NetworkTCP)
	case "UDP":
		builder = builder.Network(network.NetworkUDP)
	case "WS", "WEBSOCKET":
		builder = builder.Network(network.NetworkWebSocket).RequestURI(cfg.Protocol)
	default:
		panic("not find config.ServerConfig.ProtocolType config item!")
	}
	instance.netServer = builder.
		Addr(fmt.Sprintf("%s:%d", cfg.Ip, cfg.Port)).
		Option(gnet.WithMulticore(cfg.Multicore)).
		Build()
	return instance
}

func (im *IMServer) Setup(f func(data *[]byte, con gnet.Conn) error) error {
	im.handlerF = f
	im.netServer.InitServer()
	return nil
}

func (im *IMServer) Run() {
	network.RegisterEventHandler(im)
	im.netServer.StartServer()
}

func (im *IMServer) Stop() {
	im.netServer.StopServer()
}

func (im *IMServer) Handler(data *[]byte, c gnet.Conn) {
	err := im.handlerF(data, c)
	if err != nil {
		return
	}
}

func (*IMServer) EventType() network.NetWorkEventType {
	return network.NetWorkEventReact
}
