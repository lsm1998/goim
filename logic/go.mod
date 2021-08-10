module logic

go 1.16

replace (
	common => ../common
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc v1.39.1 => google.golang.org/grpc v1.26.0
	protocols => ../protocols
	utils => ../utils
)

require (
	common v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.3
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.39.1
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
	protocols v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
