package hello

import (
	"flag"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"testing"
	"tests/hello/proto"
	"time"
)

//etcd解析器
type etcdResolver struct {
	etcdAddr   string
	clientConn resolver.ClientConn
}

//初始化一个etcd解析器
func newResolver(etcdAddr string) resolver.Builder {
	return &etcdResolver{etcdAddr: etcdAddr}
}

func (r *etcdResolver) Scheme() string {
	return schema
}

//watch有变化以后会调用
func (r *etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
	fmt.Println(rn)
}

//解析器关闭时调用
func (r *etcdResolver) Close() {
	log.Println("Close")
}

//构建解析器 grpc.Dial()同步调用
func (r *etcdResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	//构建etcd client
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
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
func (r *etcdResolver) watch(keyPrefix string) {
	//初始化服务地址列表
	var addrList []resolver.Address

	resp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("获取服务地址列表失败：", err)
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), keyPrefix)})
		}
	}

	r.clientConn.NewAddress(addrList)

	//监听服务地址列表的变化
	rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.clientConn.NewAddress(addrList)
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					r.clientConn.NewAddress(addrList)
				}
			}
		}
	}
}

func exists(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func TestClient(t *testing.T) {
	flag.Parse()

	//注册etcd解析器
	r := newResolver(*EtcdAddr)
	resolver.Register(r)

	//客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
	conn, err := grpc.Dial(r.Scheme()+"://author/"+*ServiceName, grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败：", err)
	}
	defer conn.Close()

	//获得grpc句柄
	c := proto.NewUserApiClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Morning 调用...")
		reply, err := c.SayHello(context.Background(), &proto.SayHelloReq{
			Name: "lsm",
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Night 响应：%d，数据：%s\n", reply.Code, reply.Data)
	}
}
