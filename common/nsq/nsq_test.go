package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"testing"
	"time"
)

type demoNsqConsumer struct {
	Name string
}

func (d *demoNsqConsumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println(d.Name, " 消费 ", string(msg.Body))
	msg.Finish()
	return nil
}

func TestRegisterMQHandler(t *testing.T) {
	host := "119.91.113.111:4150"
	topic := "demo-topic"
	// 注册3个消费者
	go func() {
		RegisterMQHandler(host, topic, "c1", &demoNsqConsumer{Name: "1"})
	}()
	go func() {
		RegisterMQHandler(host, topic, "c2", &demoNsqConsumer{Name: "2"})
	}()
	go func() {
		RegisterMQHandler(host, topic, "c2", &demoNsqConsumer{Name: "3"})
	}()

	client, err := NewNsqClient(host)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(3 * time.Second)
		client.Producer(topic, []byte("hello"))
	}
}
