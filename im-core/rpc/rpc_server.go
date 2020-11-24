package rpc

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

type ImRpcServer int

func Init() {
	server.UsePool = true
	newServer := server.NewServer()
	addRegistryPlugin(newServer)
	_ = newServer.RegisterName("IM", new(ImRpcServer), "")
	_ = newServer.Serve("tcp", fmt.Sprintf(`:%d`, 12000))
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: fmt.Sprintf(`tcp@:%d`, 12000),
		ConsulServers:  []string{"127.0.0.1:8500"},
		BasePath:       "IM",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		panic(err)
	}
	s.Plugins.Add(r)
}
