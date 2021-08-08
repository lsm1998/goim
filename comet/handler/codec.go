package handler

import (
	"bytes"
	"common/net"
	"github.com/panjf2000/gnet"
)

const HeadSize = 4

type LengthCodec struct {
}

func (*LengthCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(net.Int32ToBytes(int32(len(buf))))
	buffer.Write(buf)
	return buffer.Bytes(), nil
}

func (*LengthCodec) Decode(c gnet.Conn) ([]byte, error) {
	n, buf := c.ReadN(HeadSize)
	if n < 4 {
		return nil, nil
	}
	bodyLen := int(net.BytesToInt32(buf))
	c.ShiftN(HeadSize)
	dataLen, data := c.ReadN(bodyLen)
	if dataLen < bodyLen {
		return nil, nil
	}
	c.ShiftN(dataLen)
	return data, nil
}
