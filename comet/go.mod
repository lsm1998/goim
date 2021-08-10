module comet

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
	github.com/garyburd/redigo v1.6.2
	github.com/golang/protobuf v1.5.2
	github.com/nsqio/go-nsq v1.0.8
	github.com/panjf2000/gnet v1.5.3
	github.com/prometheus/common v0.26.0
	github.com/spf13/cast v1.4.0
	github.com/stretchr/testify v1.7.0
	protocols v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
