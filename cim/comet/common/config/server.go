package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Server struct {
	Profile string
}

func InitServer(cfg *viper.Viper) *Server {
	return &Server{
		Profile: strings.ToUpper(cfg.GetString("profile")),
	}
}

var ServerConfig *Server
