package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"protocols/message"
	"strconv"
	"strings"
)

func main() {
	newClient(":9000")
}

func newClient(address string) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go func() {
		for {
			temp := make([]byte, 1024*10)
			le, err := conn.Read(temp)
			if err != nil {
				fmt.Println("连接退出，读取失败")
				break
			}
			if le > 0 {
				var rsp message.MessageRequest
				_ = proto.Unmarshal(temp[0:le], &rsp)
				switch pack := rsp.Pack.(type) {
				case *message.MessageRequest_Response:
					fmt.Println("收到回复=", string(pack.Response.Body))
				case *message.MessageRequest_Message:
					fmt.Println("收到消息=", string(pack.Message.Body))
				}
			}
		}
	}()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		log.Println("请开始下一步操作...")
		if input, err := inputReader.ReadString('\n'); err != nil {
			panic(err)
		} else {
			split := strings.Split(strings.TrimSuffix(input, "\r\n"), "-")
			if len(split) == 3 {
				var req message.MessageRequest
				req.Type = message.RequestType_Request
				msg := new(message.Message)
				if split[0] == "握手" || split[0] == "ws" {
					req.Cmd = message.MessageType_Handshake
					fmt.Println(split[2])
					msg.Body = []byte(split[2])
				} else if split[0] == "私信" || split[0] == "sx" {
					req.Cmd = message.MessageType_PrivateMessage
					msg.ToId, err = strconv.ParseInt(split[1], 10, 64)
					msg.Body = []byte(split[2])
				} else {
					log.Println("不能识别对应的标识")
					continue
				}
				req.Pack = &message.MessageRequest_Message{
					Message: msg,
				}
				bytes, _ := proto.Marshal(&req)
				conn.Write(bytes)
			} else {
				log.Println("消息格式错误")
			}
		}
	}
}
