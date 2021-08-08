package etcd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewEtcdClient(t *testing.T) {
	etcdAddr := "119.91.113.111:12379;119.91.113.111:22379;119.91.113.111:32379"
	go setEtcd(t, etcdAddr, "aaa")
	go setEtcd(t, etcdAddr, "bbb")
	go setEtcd(t, etcdAddr, "ccc")

	go getEtcd(t, etcdAddr)
	select {}
}

func setEtcd(t *testing.T, etcdAddr, serverAddr string) {
	client, err := NewEtcdClient(etcdAddr)
	assert.ErrorIs(t, err, nil)
	err = client.Register("hello", serverAddr)
	assert.ErrorIs(t, err, nil)
	select {}
}

func getEtcd(t *testing.T, etcdAddr string) {
	client, err := NewEtcdClient(etcdAddr)
	assert.ErrorIs(t, err, nil)

	err = client.Watch("hello")
	assert.ErrorIs(t, err, nil)

	for {
		time.Sleep(2 * time.Second)
		fmt.Println("list =>", client.AddrList("hello"))
		addr, err := client.GetAddr("hello")
		if err == errAddrListEmpty {
			fmt.Println("无人注册！！！")
			continue
		}
		fmt.Println("addr =>", addr)
	}
}
