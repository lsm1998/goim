package route

import (
	"github.com/panjf2000/gnet"
	"sync"
	"time"
)

const blockSize = 64

type ConnMaps []sync.Map

var connMaps ConnMaps

type connContent struct {
	c        gnet.Conn
	lastTime int64
}

func init() {
	connMaps = make([]sync.Map, blockSize)
	for i := 0; i < blockSize; i++ {
		connMaps[i] = sync.Map{}
	}
}

func (m *ConnMaps) Join(uid int64, c gnet.Conn) {
	connMaps[uid%blockSize].Store(uid, &connContent{c: c, lastTime: time.Now().Unix()})
}

func (m *ConnMaps) Leave(uid int64) {
	connMaps[uid%blockSize].Delete(uid)
}

func (m *ConnMaps) GetConn(uid int64) gnet.Conn {
	load, ok := connMaps[uid%blockSize].Load(uid)
	if !ok {
		return nil
	}
	return load.(*connContent).c
}

func (m *ConnMaps) LeaveByConn(c gnet.Conn) {
	var flag bool
	for i := 0; i < blockSize; i++ {
		connMaps[i].Range(func(key, value interface{}) bool {
			if value.(*connContent).c == c {
				flag = true
				connMaps[i].Delete(key)
				return false
			}
			return true
		})
		if flag {
			break
		}
	}
}

func (m *ConnMaps) setPong(uid int64) bool {
	content, ok := connMaps[uid%blockSize].Load(uid)
	if !ok {
		return false
	}
	content.(*connContent).lastTime = time.Now().Unix()
	return true
}

func (m *ConnMaps) Foreach(f func(uid int64, c gnet.Conn)) {
	for i := 0; i < blockSize; i++ {
		connMaps[i].Range(func(key, value interface{}) bool {
			f(key.(int64), value.(*connContent).c)
			return true
		})
	}
}
