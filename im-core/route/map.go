package route

import (
	"github.com/panjf2000/gnet"
	"utils"
)

type Connect struct {
	Conn     gnet.Conn
	PongTime int64
	AesKey   string
}

// userMap 在线列表
var userMap = new(utils.SyncMap)

// groupMap 群组列表
var groupMap = new(utils.SyncMap)
