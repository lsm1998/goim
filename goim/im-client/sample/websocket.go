package sample

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

func InitWebsocket() {
	url := "ws://" + Conf.WebsocketAddr + "/sub"
	conn, rsp, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Errorf("websocket.Dial(\"%s\") error(%v)", Conf.WebsocketAddr, err)
		return
	}
	defer rsp.Body.Close()

	fmt.Println(rsp.Status)

	data := authBytes(100, "live://1000", "web", []int{1000, 1001, 1002})
	if err = conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		fmt.Errorf("Write error(%v)", err)
		return
	}

	go func() {
		for {
			heartbeat := []byte{0, 0, 0, 16, 0, 16, 0, 1, 0, 0, 0, 2, 0, 0, 0, 1}
			fmt.Println("heart beat")
			if err := conn.WriteMessage(websocket.BinaryMessage, heartbeat); err != nil {
				panic(err)
			}
			time.Sleep(5 * time.Second)
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

func authBytes(mid uint, roomId, platform string, accepts []int) []byte {
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf(`{"mid":%d, "room_id":"%s", "platform":"%s", "accepts":[`, mid, roomId, platform))
	for i := 0; i < len(accepts); i++ {
		b.WriteString(strconv.Itoa(accepts[i]))
		if i < len(accepts)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("]}")
	bodyBuf := b.Bytes()
	buff := bytes.NewBuffer([]byte{})
	// head
	n := int(rawHeaderLen) + len(bodyBuf)
	buff.Write(Int32ToBytes(int32(n)))
	buff.Write(Int16ToBytes(rawHeaderLen))
	buff.Write(Int16ToBytes(1))
	buff.Write(Int32ToBytes(7))
	buff.Write(Int32ToBytes(1))
	// body
	buff.Write(bodyBuf)
	return buff.Bytes()
}

func Int32ToBytes(n int32) []byte {
	byteBuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(byteBuf, binary.BigEndian, n)
	return byteBuf.Bytes()
}

func Int16ToBytes(n int16) []byte {
	byteBuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(byteBuf, binary.BigEndian, n)
	return byteBuf.Bytes()
}
