package tests

import (
	"bufio"
	commonNet "common/net"
	"context"
	"encoding/json"
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	proto "protocols/message"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	// 请求登录接口
	client := http.Client{}
	body := `{"username":"123","password":"123"}`
	req, err := http.NewRequestWithContext(context.TODO(), "POST", "http://localhost:8000/user/login", strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")
	rsp, err := client.Do(req)
	assert.Nil(t, err)
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)
	var result = struct {
		Code int32
		Data string
		Msg  string
	}{}
	err = json.Unmarshal(bytes, &result)
	assert.Nil(t, err)

	if result.Code != 200 {
		t.Error("获取地址错误")
	}

	conn, err := net.Dial("tcp", result.Data)
	reader := bufio.NewReader(conn)
	go func() {
		for {
			b := make([]byte, 0, 4)
			_, err = reader.Read(b)
			assert.Nil(t, err)
			b = make([]byte, 0, commonNet.BytesToInt32(b))
			_, err = reader.Read(b)
			assert.Nil(t, err)
			handlerRecData(b)
		}
	}()
	b, err := getHandshakeData()
	assert.Nil(t, err)
	conn.Write(b)
}

func handlerRecData(b []byte) error {
	request := &proto.MessageRequest{}
	err := protobuf.Unmarshal(b, request)
	if err != nil {
		return err
	}
	fmt.Println("Pack=", request.Pack)
	return nil
}

func getHandshakeData() ([]byte, error) {
	request := &proto.MessageRequest{
		Cmd:  proto.MessageType_Handshake,
		Type: proto.RequestType_Request,
		Pack: &proto.MessageRequest_Message{Message: &proto.Message{
			FormId: 1,
			Body:   []byte("1"),
		}},
	}
	return protobuf.Marshal(request)
}
