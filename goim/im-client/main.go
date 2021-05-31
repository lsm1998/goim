package main

import (
	"client/sample"
	"flag"
	"runtime"
)

func main() {
	flag.Parse()

	sample.Conf = sample.NewConfig()
	runtime.GOMAXPROCS(2)
	//if Conf.Type == ProtoTCP {
	//sample.InitTCP()
	sample.InitWebsocket()
	//} else if Conf.Type == ProtoWebsocket {
	//	initWebsocket()
	//} else if Conf.Type == ProtoWebsocketTLS {
	//	initWebsocketTLS()
	//}
}
