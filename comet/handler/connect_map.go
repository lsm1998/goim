package handler

import (
	"github.com/panjf2000/gnet"
	"sync"
)

const blockSize = 64

type ConnMaps []map[int64]gnet.Conn

var connMaps ConnMaps

var mutexGroup []*sync.Mutex

func init() {
	mutexGroup = make([]*sync.Mutex, blockSize)
	connMaps = make([]map[int64]gnet.Conn, blockSize)
	for i := 0; i < blockSize; i++ {
		mutexGroup[i] = &sync.Mutex{}
		connMaps[i] = make(map[int64]gnet.Conn)
	}
}

func (m *ConnMaps) Join(uid int64, c gnet.Conn) {
	index := uid % blockSize
	mutex := mutexGroup[index]
	mutex.Lock()
	connMaps[index][uid] = c
	defer mutex.Unlock()
}

func (m *ConnMaps) Leave(uid int64) {
	index := uid % blockSize
	mutex := mutexGroup[index]
	mutex.Lock()
	delete(connMaps[index], uid)
	defer mutex.Unlock()
}

func (m *ConnMaps) GetConn(uid int64) gnet.Conn {
	return connMaps[uid%blockSize][uid]
}

func (m *ConnMaps) LeaveByConn(c gnet.Conn) {
	for i := 0; i < blockSize; i++ {
		m2 := connMaps[i]
		for k := range m2 {
			if m2[k] == c {
				m.Leave(k)
				return
			}
		}
	}
}
