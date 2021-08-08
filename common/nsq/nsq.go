package nsq

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var errProducerNil = errors.New("producer is nil")

type NsqClient struct {
	nsqHost  string
	producer *nsq.Producer
}

func NewNsqClient(nsqHost string) (*NsqClient, error) {
	p, err := nsq.NewProducer(nsqHost, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	return &NsqClient{
		producer: p,
	}, nil
}

// Producer nsq发布消息
func (c *NsqClient) Producer(topic string, msgBody []byte) error {
	if c.producer == nil {
		return errProducerNil
	}
	return c.producer.Publish(topic, msgBody)
}

// NsqConsumer nsq订阅消息
type NsqConsumer interface {
	HandleMessage(msg *nsq.Message) error
}

type consumerMgr struct {
	c              *nsq.Consumer
	topic, channel string
}

var mgrList []*consumerMgr

// RegisterMQHandler 注册业务handler, channel 填业务在consul注册的名称
func RegisterMQHandler(nsqHost, topic, channel string, h NsqConsumer) (err error) {
	config := nsq.NewConfig()
	config.DefaultRequeueDelay = 0
	config.MaxBackoffDuration = 20 * time.Millisecond
	config.LookupdPollInterval = 1000 * time.Millisecond
	config.RDYRedistributeInterval = 1000 * time.Millisecond
	config.MaxInFlight = 2500
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(fmt.Sprintf("RegisterMQHandler err: %v, topic: %s, channel: %s", err, topic, channel))
	}
	consumer.AddHandler(h)
	err = consumer.ConnectToNSQD(nsqHost)
	if err != nil {
		panic(fmt.Sprintf("RegisterMQHandler err: %v", err))
	}
	mgrList = append(mgrList, &consumerMgr{c: consumer, topic: topic, channel: channel})
	return
}

// Shutdown 停止消费, 在main函数结束时调用
func Shutdown() {
	for _, item := range mgrList {
		item.c.Stop()
	}
}
