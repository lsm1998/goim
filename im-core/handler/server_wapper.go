package handler

import (
	"github.com/panjf2000/gnet"
	"log"
	"time"
)

func (t *ImServer) Tick() (delay time.Duration, action gnet.Action) {
	log.Println("It's time to push data to clients!!!")
	//ps.connectedSockets.Range(func(key, value interface{}) bool {
	//	addr := key.(string)
	//	c := value.(gnet.Conn)
	//	c.AsyncWrite([]byte(fmt.Sprintf("heart beating to %s\n", addr)))
	//	return true
	//})
	var interval time.Duration

	interval = time.Second

	delay = interval
	return
}

// 服务端完成初始化后调用
func (t *ImServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("服务端监听启动完毕 on %s (multi-cores: %t, loops: %d) \n", srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

// 连接打开
func (t *ImServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

// 连接关闭
func (t *ImServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	closeWork(c)
	return
}

// 消息回应
func (t *ImServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, frame...)
	// 把任务丢到协程池异步处理
	_ = t.pool.Submit(func() {
		reactWork(data, c)
	})
	return
}
