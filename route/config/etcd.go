package config

import (
	"common"
	"strings"
)

var EtcdClient *common.EtcdClient

func initEtcd() {
	var err error
	EtcdClient, err = common.NewEtcdClient(strings.Join(C.Adders, ";"))
	if err != nil {
		panic(err)
	}
}
