module web

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/smallnest/rpcx v0.0.0-20201019093943-4119dd02e20f
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/grpc/examples v0.0.0-20201023201832-a80e65018272 // indirect
	utils v0.0.0-00010101000000-000000000000
)

replace utils => ../utils
