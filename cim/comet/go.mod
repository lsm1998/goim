module comet

go 1.16

require (
	github.com/panjf2000/gnet v1.4.5
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.0
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
