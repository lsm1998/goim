package tests

import (
	"fmt"
	"net"
	"testing"
)

func TestTCP(t *testing.T) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerHost, ServerPort))
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		conn.Write([]byte("hello"))
	}
}
