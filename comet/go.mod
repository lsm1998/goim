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
	github.com/panjf2000/gnet v1.5.3
	protocols v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
