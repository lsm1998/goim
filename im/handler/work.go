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
	switch pack := req.Pack.(type) {
	case *message.MessageRequest_Message:
		switch pack.Message.Cmd {
		case message.MessageType_Pong: // pong
			route.SetPongTime(pack.Message.FormId)
		case message.MessageType_Handshake: // 握手
			var rsp message.Reply
			if id := handshakeReq(pack.Message.Body); id > 0 {
				aseKey, _ := logic.GetAndSaveAesKey(id)
				route.Join(id, c, aseKey)
				rsp.Code = 200
				rsp.Body = []byte(aseKey)
			} else {
				rsp.Code = 201
			}
			_ = replyMessage(c, &rsp)
		case message.MessageType_Online: // 上线
		case message.MessageType_Offline: // 下线
		case message.MessageType_File: // 文件传输
		case message.MessageType_GroupMessage: // 群发消息
		case message.MessageType_PrivateMessage: // 私聊消息
			var rsp message.Reply
			conn, _, _ := route.Get(pack.Message.ToId)
			// 不在线
			if conn == nil {
				rsp.Code = 201
			} else {
				rsp.Code = 200
				_ = conn.AsyncWrite(bytes)
			}
			_ = replyMessage(c, &rsp)
		case message.MessageType_SystemBroadcast: // 系统广播
		}
	case *message.MessageRequest_Response:
		fmt.Println(pack.Response.Body)
	}
}

func closeWork(conn gnet.Conn) {
	route.ForEach(func(id int64, c *route.Connect) bool {
		if conn == c.Conn {
			route.Remove(id)
			return false
		}
		return true
	})
}
