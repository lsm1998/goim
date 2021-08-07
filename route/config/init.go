package config

import "utils"

var C *Config

type Config struct {
	Registry
	Rpc
	Redis
}

type Registry struct {
	Adders []string
}

type Rpc struct {
	Port     uint
	Server   string
	Metadata string
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
}
