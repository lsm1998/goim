module im-logic

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/bilibili/discovery v1.2.0
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/zhenjl/cityhash v0.0.0-20131128155616-cdd6a94144ab
	google.golang.org/grpc v1.38.0
	gopkg.in/Shopify/sarama.v1 v1.20.1
	im-common v0.0.0-00010101000000-000000000000
)

replace im-common => ./../im-common
