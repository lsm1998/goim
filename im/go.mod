module im

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/panjf2000/gnet v1.3.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/smallnest/rpcx v0.0.0-20200924044220-f2cdd4dea15a
	protocols v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.29.0
	protocols => ../protocols
	utils => ../utils
)
