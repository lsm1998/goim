package strategy

import (
	protobuf "github.com/golang/protobuf/proto"
	proto "protocols/message"
)

// // 心跳
//	MessageType_Pong MessageType = 0
//	// 握手
//	MessageType_Handshake MessageType = 1
//	// 上线
//	MessageType_Online MessageType = 2
//	// 下线
//	MessageType_Offline MessageType = 3
//	// 文件传输
//	MessageType_File MessageType = 4
//	// 群发消息
//	MessageType_GroupMessage MessageType = 5
//	// 私聊消息
//	MessageType_PrivateMessage MessageType = 6
//	// 系统广播
//	MessageType_SystemBroadcast MessageType = 7

var StrategyMap map[proto.MessageType]MsgHandlerStrategy

func init() {
	StrategyMap = make(map[proto.MessageType]MsgHandlerStrategy)
	StrategyMap[proto.MessageType_Pong] = &Pong{}
	StrategyMap[proto.MessageType_Handshake] = &Handshake{}
	StrategyMap[proto.MessageType_SystemBroadcast] = &Broadcast{}
}

func createOkResponse(cmd proto.MessageType) []byte {
	return createResponse(cmd, &proto.Reply{Code: 200})
}

func createFailResponse(cmd proto.MessageType, code int32, msg string) []byte {
	return createResponse(cmd, &proto.Reply{Code: code, Body: []byte(msg)})
}

func createResponse(cmd proto.MessageType, pack *proto.Reply) []byte {
	bytes, _ := protobuf.Marshal(&proto.MessageRequest{
		Type: proto.RequestType_Response,
		Cmd:  cmd,
		Pack: &proto.MessageRequest_Response{
			Response: pack,
		},
	})
	return bytes
}
