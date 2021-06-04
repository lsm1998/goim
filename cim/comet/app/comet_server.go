package app

import (
	"comet/app/handler"
	"comet/app/protocol"
	"comet/app/server"
	"comet/app/sink"
	"sync"
	"utils/chain"
)

type CometServer struct {
	event    chain.EventChain
	protocol chain.ProtocolChain
	sink     chain.SinkChain
	network  chain.UDPManager

	//同步处理
	group sync.WaitGroup
}

func (obj *CometServer) Setup() (err error) {
	server.Register()
	protocol.Register()
	handler.Register()
	sink.Register()
	if err = obj.network.Setup(); err != nil {
		return
	}
	if err = obj.protocol.Setup(); err != nil {
		return
	}
	if err = obj.event.Setup(); err != nil {
		return
	}
	if err = obj.sink.Setup(); err != nil {
		return
	}
	// 业务链串连
	obj.event.SetSink(&obj.sink)
	obj.protocol.SetSink(&obj.event)
	obj.network.SetSink(&obj.protocol)
	return nil
}

func (obj *CometServer) Run() {
	go func() {
		obj.group.Add(1)
		obj.sink.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		obj.event.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		obj.protocol.Run()
		obj.group.Done()
	}()
	go func() {
		obj.group.Add(1)
		obj.network.Run()
		obj.group.Done()
	}()
}

func (obj *CometServer) UnInit() {

}

func (obj *CometServer) Stop() {
	obj.network.Stop()
	obj.group.Wait()
}
