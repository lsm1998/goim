package rpc

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"im/config"
	"time"
)

type ImRpcServer int

func Init() {
	server.UsePool = true
	newServer := server.NewServer()
	addRegistryPlugin(newServer)
	_ = newServer.RegisterName(config.C.Rpc.Server, new(ImRpcServer), config.C.Rpc.Metadata)
	_ = newServer.Serve("tcp", fmt.Sprintf(`:%d`, config.C.Rpc.Port))
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: fmt.Sprintf(`tcp@:%d`, config.C.Rpc.Port),
		ConsulServers:  config.C.Registry.Adders,
		BasePath:       config.C.Rpc.Server,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		panic(err)
	}
	s.Plugins.Add(r)
}
