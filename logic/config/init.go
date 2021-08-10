package config

import "utils"

var C *Config

type Config struct {
	Registry
	Client
	Http
	Nsq
	Redis
	Mysql
}

type Mysql struct {
	Host     string
	User     string
	Password string
	DB       string `json:"db" yaml:"db"`
}

type Nsq struct {
	Host           string
	BroadcastTopic string `json:"broadcast_topic" yaml:"broadcast_topic"`
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

type Redis struct {
	Adder string
	Db    uint8
	Auth  string
}

func init() {
	C = new(Config)
	if err := utils.ScanConfig(C); err != nil {
		panic(err)
	}
	initOrm()
}
