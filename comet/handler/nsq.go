package handler

import "github.com/nsqio/go-nsq"
import protobuf "github.com/golang/protobuf/proto"
import proto "protocols/message"

type MessageConsumer struct {
}

func (*MessageConsumer) HandleMessage(msg *nsq.Message) error {
	var req proto.MessageRequest
	err := protobuf.Unmarshal(msg.Body, &req)
	if err != nil {
		return err
	}
	if req.Cmd == proto.MessageType_SystemBroadcast || req.Cmd == proto.MessageType_PrivateMessage {
		processorMessage(req.Cmd, req.Pack.(*proto.MessageRequest_Message), nil)
	} else {

	}
	msg.Finish()
	return nil
}
