package dao

import (
	"context"
	"flag"
	"fmt"
	"gopkg.in/Shopify/sarama.v1"
	"os"
	"testing"

	"im-common/internal/logic/conf"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	if err := flag.Set("conf", "../../../cmd/logic/logic-example.toml"); err != nil {
		panic(err)
	}
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	if err := d.Ping(context.TODO()); err != nil {
		os.Exit(-1)
	}
	if err := d.Close(); err != nil {
		os.Exit(-1)
	}
	if err := d.Ping(context.TODO()); err == nil {
		os.Exit(-1)
	}
	d = New(conf.Conf)
	os.Exit(m.Run())
}

func TestKafka(t *testing.T) {
	var c conf.Config
	c.Kafka = new(conf.Kafka)
	c.Kafka.Topic = "goim-push-topic-demo"
	c.Kafka.Brokers = []string{"119.29.117.244:9092"}

	c.Redis = new(conf.Redis)
	c.Redis.Network = "tcp"
	c.Redis.Addr = "192.168.74.128:6379"
	dao := New(&c)
	err := dao.BroadcastMsg(context.Background(), 1, 1, []byte{1, 2, 3, 4, 5})
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = `nginx_log`
	msg.Value = sarama.StringEncoder("this is a good test")
	client, err := sarama.NewSyncProducer([]string{"119.29.117.244:9092"}, config)
	if err != nil {
		fmt.Println("producer close err, ", err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)

	if err != nil {
		fmt.Println("send message failed, ", err)
		return
	}
	fmt.Printf("分区ID:%v, offset:%v \n", pid, offset)
}
