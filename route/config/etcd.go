package config

import (
	"common/etcd"
	"strings"
)

var EtcdClient *etcd.EtcdClient

func initEtcd() {
	var err error
	EtcdClient, err = etcd.NewEtcdClient(strings.Join(C.Adders, ";"))
	if err != nil {
		panic(err)
	}
}
