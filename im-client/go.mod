module client

go 1.15

require (
	github.com/golang/protobuf v1.4.3
	github.com/panjf2000/gnet v1.3.0
	protocols v0.0.0-00010101000000-000000000000
)

replace (
	protocols => ../protocols
)
