package build

import (
	"comet"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/gnet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWSNetServer(t *testing.T) {
	comet.RegisterEventHandler(new(echoDemo))
	builder := &NetServerBuilder{}
	server := builder.Addr(":8500").
		Network(comet.NetworkWebSocket).
		RequestURI("/sub").
		Option(gnet.WithMulticore(true)).
		Build()
	server.StartServer()
}

func TestWSClient(t *testing.T) {
	url := "ws://127.0.0.1:8500/sub"
	conn, rsp, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
	defer rsp.Body.Close()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			err = conn.WriteMessage(websocket.BinaryMessage, []byte("hello"))
			assert.NoError(t, err)
		}
	}()
	for {
		fmt.Println("recv--")
		_, recvBuf, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(recvBuf)
	}
}
