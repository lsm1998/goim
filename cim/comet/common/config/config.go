package config

import (
	"comet/common/filesystem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func Setup(path string) error {
	if filesystem.FileExist(path) {
		viper.SetConfigFile(path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
			return err
		}

		//Replace environment variables
		err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
		if err != nil {
			log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
			return err
		}
	} else {
		up, err := url.Parse(path)
		if err != nil {
			log.Fatal(fmt.Sprintf("remote config fail: %s", err.Error()))
			return err
		}
		err = viper.AddRemoteProvider(up.Scheme, up.Host, up.Path)
		if err != nil {
			log.Fatal(fmt.Sprintf("remote config fail: %s", err.Error()))
			return err
		}
		viper.SetConfigType("yaml")    //以yaml格式进行读取
		err = viper.ReadRemoteConfig() //正式读取
		if err != nil {
			log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
			return err
		}
	}
	server := viper.Sub("settings.server")
	if server == nil {
		panic("No found settings.server in the configuration")
	}
	ServerConfig = InitServer(server)

	network := viper.Sub("settings.network")
	if network == nil {
		panic("No found settings.udp in the configuration")
	}
	NetworkConfig = InitNetwork(network)

	//logger := viper.Sub("settings.logger")
	//if logger == nil {
	//	panic("No found settings.logger in the configuration")
	//}
	//LoggerConfig = InitLogger(logger)
	return nil
}
