package config

import "utils"

var C *Config

type Config struct {
	Registry
	Server
	Nsq
	Redis
}

type Registry struct {
	Adders []string
}

type Nsq struct {
	Host string
}

type Server struct {
	Port      uint16
	Name      string
	Zone      uint8
	Multicore bool
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
	initEtcd()
	initRedis()
}
