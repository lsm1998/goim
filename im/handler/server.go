package handler

import (
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type ImServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (t *ImServer) Init() {
	t.pool = goroutine.Default()
}

func (t *ImServer) Stop() {
	t.pool.Release()
}
