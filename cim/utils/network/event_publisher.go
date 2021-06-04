package network

import (
	"github.com/panjf2000/gnet"
)

type EventPublisher struct {
	*gnet.EventServer
}

func (*EventPublisher) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	EventNotify(NetWorkEventReact, &frame, c)
	return
}

func (*EventPublisher) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	EventNotify(NetWorkEventOpen, nil, c)
	return
}

func (*EventPublisher) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	EventNotify(NetWorkEventClose, nil, c)
	return
}
