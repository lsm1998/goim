package config

import "utils"

var C *Config

type Config struct {
	Registry
	Client
	Http
	Nsq
}

type Nsq struct {
	Host string
}

type Http struct {
	Port uint16
}

type Registry struct {
	Adders []string
}

type Client struct {
	Servers []string
}

func init() {
	C = new(Config)
	if err := utils.ScanConfig(C); err != nil {
		panic(err)
	}
}
