package tests

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
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
	req.Header.Set("Content-Type", "application/json")
	rsp, err := client.Do(req)
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(bytes), err)
}
