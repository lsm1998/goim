package nsq

import (
	"common/nsq"
	"logic/config"
)

var nsqClient *nsq.NsqClient

func init() {
	var err error
	nsqClient, err = nsq.NewNsqClient(config.C.Nsq.Host)
	if err != nil {
		panic(err)
	}
}

func Producer(topic string, msgBody []byte) error {
	return nsqClient.Producer(topic, msgBody)
}
