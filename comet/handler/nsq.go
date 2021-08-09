package handler

import "github.com/nsqio/go-nsq"

type MessageConsumer struct {
}

func (*MessageConsumer) HandleMessage(msg *nsq.Message) error {
	msg.Finish()
	return nil
}
