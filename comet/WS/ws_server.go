package WS

import (
	"comet"
	"comet/WS/bufio"
	"github.com/panjf2000/gnet"
	"net"
)

type wsNetworkServer struct {
	addr       string
	opts       []gnet.Option
	requestURI string
}

func (*wsNetworkServer) InitServer() {

}

func (*wsNetworkServer) StopServer() {
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
		comet.EventNotify(comet.NetWorkEventOpen, nil, ws)
		go w.listen(ws)
	}
}

func NewWsNetworkServer(addr, requestURI string, opts ...gnet.Option) *wsNetworkServer {
	server := new(wsNetworkServer)
	server.addr = addr
	server.opts = opts
	server.requestURI = requestURI
	return server
}

func (w *wsNetworkServer) listen(ws *WSConn) {
	for {
		op, data, err := ws.ReadMessage()
		if err != nil || op != BinaryMessage {
			comet.EventNotify(comet.NetWorkEventClose, nil, ws)
			break
		}
		comet.EventNotify(comet.NetWorkEventReact, &data, ws)
	}
}
