package etcd

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"sync"
	"time"
)

const ttlDuration = 5

const schema = "ns"

var errClientNil = errors.New("clientv3.Client is nil")

var errAddrListEmpty = errors.New("addr list is empty")

type EtcdClient struct {
	etcdAddr    string
	mutex       sync.Mutex
	cli         *clientv3.Client
	registerMap map[string][]string
	count       int64
}

func NewEtcdClient(etcdAddr string) (*EtcdClient, error) {
	return &EtcdClient{
		etcdAddr:    etcdAddr,
		mutex:       sync.Mutex{},
		registerMap: map[string][]string{},
	}, nil
}

// Register 注册服务
func (c *EtcdClient) Register(serviceName, serverAddr string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cli == nil {
		if err := c.initClient(); err != nil {
			return err
		}
	}
	//与etcd建立长连接，并保证连接不断(心跳检测)
	ticker := time.NewTicker(time.Second * time.Duration(ttlDuration))
	go func() {
		key := "/" + schema + "/" + serviceName + "/" + serverAddr
		for {
			resp, err := c.cli.Get(context.Background(), key)
			if err != nil {
				fmt.Printf("获取服务地址失败：%s", err)
			} else if resp.Count == 0 {
				// 尚未注册
				err = c.keepAlive(serviceName, serverAddr, ttlDuration)
				if err != nil {
					fmt.Printf("保持连接失败：%s", err)
				}
			}
			<-ticker.C
		}
	}()
	return nil
}

// UnRegister 服务退出
func (c *EtcdClient) UnRegister(serviceName, serverAddr string) error {
	var err error
	if c.cli != nil {
		key := "/" + schema + "/" + serviceName + "/" + serverAddr
		_, err = c.cli.Delete(context.Background(), key)
	} else {
		err = errClientNil
	}
	return err
}

// Watch 监听
func (c *EtcdClient) Watch(serviceName string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cli == nil {
		if err := c.initClient(); err != nil {
			return err
		}
	}
	go func() {
		keyPrefix := "/" + schema + "/" + serviceName + "/"
		//初始化服务地址列表
		var addrList = c.registerMap[serviceName]
		//监听服务地址列表的变化
		rch := c.cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
		for n := range rch {
			for _, ev := range n.Events {
				addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
				switch ev.Type {
				case mvccpb.PUT:
					if !exists(addrList, addr) {
						c.registerMap[serviceName] = append(c.registerMap[serviceName], addr)
					}
				case mvccpb.DELETE:
					if s, ok := remove(addrList, addr); ok {
						c.registerMap[serviceName] = s
					}
				}
			}
		}
	}()
	return nil
}

// GetAddr 获取服务对应的一个地址
func (c *EtcdClient) GetAddr(serviceName string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	list := c.registerMap[serviceName]
	if len(list) == 0 {
		return "", errAddrListEmpty
	}
	c.count++
	return list[c.count%int64(len(list))], nil
}

// AddrList 获取服务对应的地址列表
func (c *EtcdClient) AddrList(serviceName string) []string {
	return c.registerMap[serviceName]
}

// PeekLive 是否存活
func (c *EtcdClient) PeekLive(serviceName, addr string) bool {
	return exists(c.registerMap[serviceName], addr)
}

func (c *EtcdClient) initClient() error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(c.etcdAddr, ";"),
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return err
	}
	c.cli = cli
	return nil
}

func (c *EtcdClient) keepAlive(serviceName, serverAddr string, ttl int64) error {
	client := c.cli
	//创建租约
	leaseResp, err := client.Grant(context.Background(), ttl)
	if err != nil {
		fmt.Printf("创建租期失败：%s\n", err)
		return err
	}
	//将服务地址注册到etcd中
	key := "/" + schema + "/" + serviceName + "/" + serverAddr
	_, err = client.Put(context.Background(), key, serverAddr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		fmt.Printf("注册服务失败：%s", err)
		return err
	}
	//建立长连接
	ch, err := client.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		fmt.Printf("建立长连接失败：%s\n", err)
		return err
	}
	// 清空keepAlive返回的channel
	go func() {
		for {
			<-ch
		}
	}()
	return nil
}

func exists(l []string, addr string) bool {
	for i := range l {
		if l[i] == addr {
			return true
		}
	}
	return false
}

func remove(s []string, addr string) ([]string, bool) {
	for i := range s {
		if s[i] == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
