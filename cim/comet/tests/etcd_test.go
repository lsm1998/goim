package tests

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"testing"
	"time"
)

const HOST = "127.0.0.1:2379"

var client *clientv3.Client

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{HOST},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to etcd success")
	client = cli
}

func TestRegNode(t *testing.T) {

}

/**
基本的 key-value 存取
*/
func put(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := client.Put(ctx, key, value)
	cancel()
	if err != nil {
		panic(err)
	}
}

/**
基本的 key-value 存取
*/
func get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		panic(err)
	}
	for _, ev := range resp.Kvs {
		return string(ev.Value), nil
	}
	return "", errors.New("not find!")
}

/**
定时的 key-value 存取
timeOut（秒）后自动执行DELETE事件
*/
func lease(key, value string, timeOut int64) {
	resp, err := client.Grant(context.TODO(), timeOut)
	if err != nil {
		panic(err)
	}
	_, err = client.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}
}

// keepAlive 延长租期
func keepAlive(key, value string) {
	resp, err := client.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	_, err = client.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}
	ch, err := client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		panic(err)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

// watch 监听键值对
func watch(key string, handle func(event *clientv3.Event)) {
	// 返回一个 WatchResponse channel
	rch := client.Watch(context.Background(), key)
	for temp := range rch {
		for _, ev := range temp.Events {
			handle(ev)
		}
	}
}
