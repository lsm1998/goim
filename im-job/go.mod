module im-job

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/bilibili/discovery v1.2.0
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.38.0
	im-common v0.0.0-00010101000000-000000000000
)

replace im-common => ../im-common
