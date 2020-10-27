package logic

import (
	"im/config"
	"im/model"
	"protocols/message"
	"time"
	"utils"
)

// 消息落库
func SaveMessage(msg *message.Message) error {
	s := &model.Message{Body: msg.Body, Cmd: uint(msg.Cmd), FormId: msg.FormId, ToId: msg.ToId}
	s.UserTime = utils.Unix2Time(msg.CreateTime)
	s.UpdateTime = time.Now().UTC()
	s.CreateTime = s.UpdateTime
	s.AesKey, _ = QueryAesKey(msg.FormId)
	return config.DB.Save(s).Error
}
