package WS

import (
	"common/net/network"
	"common/net/network/WS/bufio"
	"github.com/panjf2000/gnet"
	"net"
)

type wsNetworkServer struct {
	addr         string
	opts         []gnet.Option
	requestURI   string
	initF, stopF func()
}

func (w *wsNetworkServer) InitServer() {
	if w.initF != nil {
		w.initF()
	}
}

func (w *wsNetworkServer) StopServer() {
	if w.stopF != nil {
		w.stopF()
	}
}

func (w *wsNetworkServer) StartServer() {
	ln, err := net.Listen("tcp", w.addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		rd := bufio.NewReader(conn)
		wr := bufio.NewWriter(conn)
		req, err := ReadRequest(rd)
		if err != nil {
			continue
		}
		if req.RequestURI != w.requestURI {
			continue
		}
		ws, err := Upgrade(conn, rd, wr, req)
		network.EventNotify(network.NetWorkEventOpen, nil, ws)
		go w.listen(ws)
	}
}

func NewWsNetworkServer(addr, requestURI string, initF, stopF func(), opts ...gnet.Option) *wsNetworkServer {
	server := new(wsNetworkServer)
	server.addr = addr
	server.opts = opts
	server.requestURI = requestURI
	server.initF, server.stopF = initF, stopF
	return server
}

func (w *wsNetworkServer) listen(ws *WSConn) {
	for {
		op, data, err := ws.ReadMessage()
		if err != nil || op != BinaryMessage {
			network.EventNotify(network.NetWorkEventClose, nil, ws)
			break
		}
		network.EventNotify(network.NetWorkEventReact, &data, ws)
	}
}
