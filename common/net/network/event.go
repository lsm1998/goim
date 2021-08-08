package network

import "github.com/panjf2000/gnet"

type NetWorkEventType uint8

const (
	NetWorkEventReact NetWorkEventType = iota + 1
	NetWorkEventOpen
	NetWorkEventClose
)

type NetWorkEvent interface {
	Handler(data []byte, c gnet.Conn)

	EventType() NetWorkEventType
}

var subscribe []NetWorkEvent

func RegisterEventHandler(event ...NetWorkEvent) {
	subscribe = append(subscribe, event...)
}

func EventNotify(typ NetWorkEventType, data []byte, c gnet.Conn) {
	for i := 0; i < len(subscribe); i++ {
		if typ == subscribe[i].EventType() {
			subscribe[i].Handler(data, c)
		}
	}
}
