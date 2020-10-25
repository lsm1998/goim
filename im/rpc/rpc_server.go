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
	_ = newServer.Serve("tcp", fmt.Sprintf(`:%d`, 11000))
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: fmt.Sprintf(`tcp@:%d`, 11000),
		ConsulServers:  []string{"47.103.211.234:8500"},
		BasePath:       "/im",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		panic(err)
	}
	s.Plugins.Add(r)
}
