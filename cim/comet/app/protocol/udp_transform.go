package protocol

import (
	"github.com/panjf2000/gnet"
	"utils/chain"
)

type UDPTransformLine struct {
}

func newUDPTransformLine() *UDPTransformLine {
	return &UDPTransformLine{}
}

func (n *UDPTransformLine) Line(data *[]byte, conn gnet.Conn) (line chain.EventLine, err error) {

	return line, nil
}
