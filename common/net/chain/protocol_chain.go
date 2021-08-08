package chain

import (
	"context"
	"errors"
	"github.com/panjf2000/gnet"
	log "github.com/sirupsen/logrus"
)

var protocolSink []TransformLine

type ProtocolChain struct {
	sink   Sink
	ctx    context.Context
	cancel context.CancelFunc
}

func (obj *ProtocolChain) SetSink(sink Sink) {
	obj.sink = sink
}

func (obj *ProtocolChain) Line(data *[]byte, conn gnet.Conn) error {
	var err error
	var eventLine EventLine
	if eventLine, err = obj.Transform(data, conn); err != nil {
		log.Errorf("Transform line fail,err=%v \n", err)
	} else if err = obj.sink.PushLine(eventLine); err != nil {
		log.Errorf("PushLine to sink fail,err=%v \n", err)
	}
	return err
}
func (obj *ProtocolChain) Setup() error {
	obj.ctx, obj.cancel = context.WithCancel(context.Background())
	return nil
}

func (obj *ProtocolChain) Run() {
	// 通知退出
	// obj.sink.Stop()
}

// Transform 根据协议版本转换EventLine
func (obj *ProtocolChain) Transform(data *[]byte, conn gnet.Conn) (line EventLine, err error) {
	for _, t := range protocolSink {
		if line, err = t.Line(data, conn); err == nil {
			return
		}
	}
	return nil, errors.New("unknown protocol")
}

func (obj *ProtocolChain) Stop() {
	obj.cancel()
	//阻塞: 外层会等待Run退出 ，所以这里可以没有wait都可以
}

func (obj *ProtocolChain) UnInit() {

}
