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
				var rsp message.MessageResponse
				_ = proto.Unmarshal(temp[0:le], &rsp)
				fmt.Println("收到回复=>", rsp)
			}
		}
	}()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		log.Println("请开始下一步操作...")
		if input, err := inputReader.ReadString('\n'); err != nil {
			panic(err)
		} else {
			split := strings.Split(strings.TrimSuffix(input, "\n"), "-")
			if len(split) == 3 {
				var req message.MessageRequest
				req.Message = new(message.Message)
				if split[0] == "握手" || split[0] == "ws" {
					req.Message.Cmd = message.RequestType_Handshake
					req.Message.Body = []byte(split[2])
				} else if split[0] == "私信" || split[0] == "sx" {
					req.Message.Cmd = message.RequestType_PrivateMessage
					req.Message.ToId, err = strconv.ParseInt(split[1], 10, 64)
					req.Message.Body = []byte(split[2])
				} else {
					log.Println("不能识别对应的标识")
					continue
				}
				bytes, _ := proto.Marshal(&req)
				conn.Write(bytes)
			} else {
				log.Println("消息格式错误")
			}
		}
	}

	//wg.Add(11)
	//for i := 0; i < 10; i++ {
	//	conn.Write([]byte("hello"))
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("发送一次")
	//	wg.Done()
	//}
	//wg.Wait()
	//conn.Close()
}
