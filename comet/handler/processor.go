package handler

import (
	"comet/handler/strategy"
	"github.com/panjf2000/gnet"
	"github.com/prometheus/common/log"
	proto "protocols/message"
)

func processorMessage(cmd proto.MessageType, message *proto.MessageRequest_Message, c gnet.Conn) {
	s, ok := strategy.StrategyMap[cmd]
	if !ok {
		log.Error("strategy type not support")
		return
	}
	err := strategy.NewMsgHandler(message.Message, c, s)
	if err != nil {
		log.Error(err)
	}
}

func processorResponse(cmd proto.MessageType, response *proto.MessageRequest_Response, c gnet.Conn) {

}
