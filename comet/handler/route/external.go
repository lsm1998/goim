package route

import (
	"comet/errors"
	"github.com/panjf2000/gnet"
)

func Pong(uid int64) error {
	if !connMaps.setPong(uid) {
		return errors.ErrNotOnlineLogin
	}
	return nil
}

func Join(uid int64, c gnet.Conn) error {
	if connMaps.GetConn(uid) != nil {
		return errors.ErrRepeatLogin
	}
	connMaps.Join(uid, c)
	// 用户上线
	return nil
}
