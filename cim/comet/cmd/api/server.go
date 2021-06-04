package api

import (
	"comet/App"
	"comet/common/config"
	"comet/common/global"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	configYml string
	server    *app.CometServer
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      global.AppName + " server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			//初始化
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			//运行
			return run()
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			// 这里做反初始化
			server.UnInit()
			server = nil
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	//1. 读取配置
	if err := config.Setup(configYml); err != nil {
		return
	}

	//2. 服务初始化
	server = new(app.CometServer)
	if err := server.Setup(); err != nil {
		log.Fatal(err)
		return
	}
	usageStr := `starting api server`
	log.Info(usageStr)
}

func run() error {
	tip()

	server.Run()

	fmt.Println("Server run at:")

	// fmt.Printf("-  Local:  %s:%v/ \r\n", config.UDPServerConfig.NetIP, config.UDPServerConfig.Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	server.Stop()
	return nil
}
func tip() {
	usageStr := `欢迎使用 ` + global.AppName + global.Version + ` 可以使用 -h 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}
