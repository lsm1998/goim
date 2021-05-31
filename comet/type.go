package comet

type Network string

const (
	NetworkTCP       Network = "tcp"
	NetworkUDP       Network = "udp"
	NetworkWebSocket Network = "ws"
)

type NetworkServer interface {
	InitServer()

	StartServer()

	StopServer()
}
