package service

import (
	protobuf "github.com/golang/protobuf/proto"
	"logic/config"
	"logic/nsq"
	proto "protocols/message"
	"time"
)

type PushService struct {
}

func NewPushService() *PushService {
	return &PushService{}
}

func (p *PushService) Broadcast(uid int64, content string) error {
	request := proto.MessageRequest{}
	request.Type = proto.RequestType_Request
	request.Cmd = proto.MessageType_SystemBroadcast
	request.Pack = &proto.MessageRequest_Message{
		Message: &proto.Message{
			FormId:     uid,
			Body:       []byte(content),
			Length:     int32(len(content)),
			CreateTime: time.Now().Unix(),
		},
	}
	bytes, err := protobuf.Marshal(&request)
	if err != nil {
		return err
	}
	if err = nsq.Producer(config.C.BroadcastTopic, bytes); err != nil {
		return err
	}
	return nil
}
