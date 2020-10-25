package handler

import "strconv"

func handshake(body []byte) int64 {
	int64, err := strconv.ParseInt(string(body), 10, 64)
	if err != nil {
		return -1
	} else {
		return int64
	}
}
