module logic-user

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/smallnest/rpcx v0.0.0-20201027145221-c31b15be63d4
	protocols v0.0.0-00010101000000-000000000000
    utils v0.0.0-00010101000000-000000000000
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.29.0
	protocols => ../protocols
	utils => ../utils
)
