package model

import (
	"time"
	"utils"
)

type MessageStatus int32

const (
	// 未读
	MsgUnread MessageStatus = iota
	// 已读
	MsgRead
	// 删除
	MsgDelete
)

type Message struct {
	utils.Model
	Cmd      uint      `json:"cmd" gorm:"cmd"`
	FormId   int64     `json:"form_id" gorm:"form_id"`
	ToId     int64     `json:"to_id" gorm:"to_id"`
	Body     []byte    `json:"body" gorm:"body"`
	AesKey   string    `json:"aes_key" gorm:"aes_key"`
	UserTime time.Time `json:"user_time" gorm:"user_time"`
}

func (*Message) TableName() string {
	return "t_message"
}

type MessageRead struct {
	utils.Model
	MessageId int64         `json:"message_id" gorm:"message_id"`
	Status    MessageStatus `json:"status" gorm:"status"`
}

func (*MessageRead) TableName() string {
	return "t_message_read"
}
