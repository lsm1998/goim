package build

import (
	"bytes"
	"common/net"
	"common/net/network"
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/stretchr/testify/assert"
	"net"
	"sync"
	"testing"
)

const HeadSize = 4

type MyCodec struct {
}

func (*MyCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(utils.Int32ToBytes(int32(len(buf))))
	buffer.Write(buf)
	return buffer.Bytes(), nil
}

func (*MyCodec) Decode(c gnet.Conn) ([]byte, error) {
	n, buf := c.ReadN(HeadSize)
	if n < 4 {
		return nil, nil
	}
	bodyLen := int(utils.BytesToInt32(buf))
	c.ShiftN(HeadSize)
	dataLen, data := c.ReadN(bodyLen)
	if dataLen < bodyLen {
		return nil, nil
	}
	c.ShiftN(dataLen)
	return data, nil
}

type echoDemo struct {
}

func (*echoDemo) Handler(data *[]byte, c gnet.Conn) {
	fmt.Println(string(*data))
	_ = c.AsyncWrite([]byte("echo:" + string(*data)))
}

func (*echoDemo) EventType() network.NetWorkEventType {
	return network.NetWorkEventReact
}

func TestNetServer(t *testing.T) {
	network.RegisterEventHandler(new(echoDemo))
	builder := &NetServerBuilder{}
	server := builder.Addr(":8500").
		Network(network.NetworkTCP).
		Option(gnet.WithMulticore(true), gnet.WithCodec(new(MyCodec))).
		Build()
	server.StartServer()
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8500")
	assert.NoError(t, err)

	go func() {
		for {
			buf := make([]byte, 1024)
			n, _ := conn.Read(buf)
			fmt.Println(string(buf[0:n]))
		}
	}()

	for i := 0; i < 10; i++ {
		b := []byte("hello")
		buffer := bytes.NewBuffer([]byte{})
		buffer.Write(utils.Int32ToBytes(int32(len(b))))
		buffer.Write(b)
		_, err = conn.Write(buffer.Bytes())
		assert.NoError(t, err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
