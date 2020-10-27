package handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/panjf2000/gnet"
	"im/logic"
	"im/route"
	"protocols/message"
)

func reactWork(bytes []byte, c gnet.Conn) {
	var req message.MessageRequest
	if err := proto.Unmarshal(bytes, &req); err != nil {
		fmt.Println("解码错误")
		return
	}
	rsp := message.MessageResponse{}
	switch {
	// Pong
	case req.Message.Cmd == message.RequestType_Pong:
		route.SetPongTime(req.Message.FormId)
	// 握手
	case req.Message.Cmd == message.RequestType_Handshake:
		if id := handshake(req.Message.Body); id > 0 {
			aseKey, _ := logic.GetAndSaveAesKey(id)
			route.Join(id, c, aseKey)
			rsp.Code = 200
			rsp.Message = "abcdef"
		} else {
			rsp.Code = 202
		}
		if data, err := proto.Marshal(&rsp); err == nil {
			_ = c.AsyncWrite(data)
		}
	// 上线
	case req.Message.Cmd == message.RequestType_Online:
	// 下线
	case req.Message.Cmd == message.RequestType_Offline:
	// 文件传输
	case req.Message.Cmd == message.RequestType_File:
	// 群发消息
	case req.Message.Cmd == message.RequestType_GroupMessage:
	// 私聊消息
	case req.Message.Cmd == message.RequestType_PrivateMessage:
		conn, _, _ := route.Get(req.Message.ToId)
		if conn == nil {
			// 不在线
			rsp.Code = 201
		} else {
			_ = conn.AsyncWrite(bytes)
			rsp.Code = 200
		}
		if data, err := proto.Marshal(&rsp); err == nil {
			_ = c.AsyncWrite(data)
		}
	// 系统广播
	case req.Message.Cmd == message.RequestType_SystemBroadcast:
	}
}

func closeWork(conn gnet.Conn) {
	route.ForEach(func(id int64, c route.Connect) bool {
		if conn == c.Conn {
			route.Remove(id)
			return false
		}
		return true
	})
}
