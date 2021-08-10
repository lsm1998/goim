package main

import (
	"comet/config"
	"comet/handler"
	"common/net/network"
	"common/net/network/build"
	"common/nsq"
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/spf13/cast"
)

func main() {
	host := fmt.Sprintf("127.0.0.1:%d", config.C.Server.Port)

	// 将服务地址注册到etcd中
	err := config.EtcdClient.Register(config.C.Server.Name, host)
	if err != nil {
		panic(err)
	}

	go func() {
		// 群聊的MQHandler
		channel := config.C.BroadcastTopic + cast.ToString(config.C.Zone)
		err = nsq.RegisterMQHandler(config.C.Nsq.Host, config.C.BroadcastTopic, channel, new(handler.MessageConsumer))
		if err != nil {
			panic(err)
		}
	}()
	//go func() {
	//	// 私聊的MQHandler
	//	channel := config.C.Topic + cast.ToString(config.C.Zone)
	//	err = nsq.RegisterMQHandler(config.C.Nsq.Host, config.C.Topic, channel, new(handler.MessageConsumer))
	//	if err != nil {
	//		panic(err)
	//	}
	//}()

	network.RegisterEventHandler(new(handler.ReactHandler), new(handler.CloseHandler))

	builder := &build.NetServerBuilder{}
	netServer := builder.
		SetUp(func() {
			fmt.Println("start!")
		}).
		Network(network.NetworkTCP).
		Addr(fmt.Sprintf("127.0.0.1:%d", config.C.Server.Port)).
		Option(
			gnet.WithMulticore(config.C.Server.Multicore),
			gnet.WithCodec(new(handler.LengthCodec)),
		).
		Build()
	netServer.InitServer()
	netServer.StartServer()
	defer nsq.Shutdown()
	defer netServer.StopServer()
}
