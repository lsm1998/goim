package route

import (
	"github.com/panjf2000/gnet"
	"sync"
	"time"
)

const blockSize = 64

type ConnMaps []map[int64]*connContent

var connMaps ConnMaps

var mutexGroup []*sync.Mutex

type connContent struct {
	c        gnet.Conn
	lastTime int64
}

func init() {
	mutexGroup = make([]*sync.Mutex, blockSize)
	connMaps = make([]map[int64]*connContent, blockSize)
	for i := 0; i < blockSize; i++ {
		mutexGroup[i] = &sync.Mutex{}
		connMaps[i] = make(map[int64]*connContent)
	}
}

func (m *ConnMaps) Join(uid int64, c gnet.Conn) {
	index := uid % blockSize
	mutex := mutexGroup[index]
	mutex.Lock()
	defer mutex.Unlock()
	connMaps[index][uid] = &connContent{c: c, lastTime: time.Now().Unix()}
}

func (m *ConnMaps) Leave(uid int64) {
	index := uid % blockSize
	mutex := mutexGroup[index]
	mutex.Lock()
	defer mutex.Unlock()
	delete(connMaps[index], uid)
}

func (m *ConnMaps) GetConn(uid int64) gnet.Conn {
	content := connMaps[uid%blockSize][uid]
	if content == nil {
		return nil
	}
	return content.c
}

func (m *ConnMaps) LeaveByConn(c gnet.Conn) {
	for i := 0; i < blockSize; i++ {
		m2 := connMaps[i]
		for k := range m2 {
			if m2[k].c == c {
				m.Leave(k)
				return
			}
		}
	}
}

func (m *ConnMaps) setPong(uid int64) bool {
	index := uid % blockSize
	mutex := mutexGroup[index]
	mutex.Lock()
	defer mutex.Unlock()
	content := connMaps[uid%blockSize][uid]
	if content == nil {
		return false
	}
	content.lastTime = time.Now().Unix()
	return true
}
