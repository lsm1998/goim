package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Network struct {
	Protocol   string
	Ip         string
	Port       uint16
	Multicore  bool
	RequestURI string
}

func InitNetwork(cfg *viper.Viper) *Network {
	return &Network{
		Protocol:   strings.ToUpper(cfg.GetString("protocol")),
		Ip:         cfg.GetString("ip"),
		Port:       uint16(cfg.GetUint("port")),
		Multicore:  cfg.GetBool("multicore"),
		RequestURI: cfg.GetString("requestURI"),
	}
}

var NetworkConfig *Network
