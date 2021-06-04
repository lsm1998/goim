package tests

import (
	"flag"
	"fmt"
	"net"
	"testing"
)

const (
	ServerHost = "10.10.30.232"
	ServerPort = 8401
)

func TestUDP(t *testing.T) {
	flag.Parse()
	fmt.Println("start,host=", ServerHost, ",port=", ServerPort)
	for i := 0; i < 100; i++ {
		sendWork()
	}
}

func sendWork() {
	svrAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ServerHost, ServerPort))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, svrAddr)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
}
