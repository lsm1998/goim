package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
	"strings"
	"time"
)

// EtcdGrpcResolver etcd解析器
type EtcdGrpcResolver struct {
	etcdAddr   string
	clientConn resolver.ClientConn
	cli        *clientv3.Client
}

// NewResolver 初始化一个etcd解析器
func NewResolver(etcdAddr string) resolver.Builder {
	return &EtcdGrpcResolver{etcdAddr: etcdAddr}
}

func (r *EtcdGrpcResolver) Scheme() string {
	return schema
}

// ResolveNow watch有变化以后会调用
func (r *EtcdGrpcResolver) ResolveNow(rn resolver.ResolveNowOptions) {

}

// Close 解析器关闭时调用
func (r *EtcdGrpcResolver) Close() {

}

// Build 构建解析器 grpc.Dial()同步调用
func (r *EtcdGrpcResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	//构建etcd client
	if r.cli == nil {
		r.cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			fmt.Printf("连接etcd失败：%s\n", err)
			return nil, err
		}
	}

	r.clientConn = clientConn

	go r.watch("/" + target.Scheme + "/" + target.Endpoint + "/")

	return r, nil
}

//监听etcd中某个key前缀的服务地址列表的变化
func (r *EtcdGrpcResolver) watch(keyPrefix string) {
	//初始化服务地址列表
	var addrList []resolver.Address

	resp, err := r.cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("获取服务地址列表失败：", err)
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), keyPrefix)})
		}
	}

	r.clientConn.NewAddress(addrList)

	//监听服务地址列表的变化
	rch := r.cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !existsAddress(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.clientConn.NewAddress(addrList)
				}
			case mvccpb.DELETE:
				if s, ok := removeAddress(addrList, addr); ok {
					addrList = s
					r.clientConn.NewAddress(addrList)
				}
			}
		}
	}
}

func existsAddress(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func removeAddress(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
