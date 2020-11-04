package handler

import (
	"utils"
)

func handshakeReq(body []byte, uid int64) int64 {
	_, err := utils.ValidToken(string(body), uid)
	if err != nil {
		return 0
	}
	return uid
}
