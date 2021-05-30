module im-comet

go 1.16

require (
	github.com/bilibili/discovery v1.2.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	im-common v0.0.0-00010101000000-000000000000
)

replace (
	im-common => ../im-common
)
