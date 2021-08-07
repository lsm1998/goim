module protocols

go 1.15

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc v1.39.1 => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.39.1
)
