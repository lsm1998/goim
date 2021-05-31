package WS

import (
	"bytes"
	"comet/common"
	"github.com/panjf2000/gnet"
)

type WebSocketCodec struct {
}

func (*WebSocketCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(common.Int32ToBytes(int32(len(buf))))
	buffer.Write(buf)
	return buffer.Bytes(), nil
}

func (*WebSocketCodec) Decode(c gnet.Conn) ([]byte, error) {

	return nil, nil
}
