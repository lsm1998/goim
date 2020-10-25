package model

import "utils"

type Message struct {
	utils.Model
	Cmd    uint   `json:"cmd" gorm:"cmd"`
	FormId int64  `json:"form_id" gorm:"form_id"`
	ToId   int64  `json:"to_id" gorm:"to_id"`
	Body   []byte `json:"body" gorm:"body"`
	Length int32  `json:"length" gorm:"length"`
	AesKey string `json:"aes_key" gorm:"aes_key"`
}

func (Message) TableName() string {
	return "t_message"
}

type MessageRead struct {
	utils.Model
	MessageId int64 `json:"message_id" gorm:"message_id"`
	Status    int32 `json:"status" gorm:"status"`
}

func (MessageRead) TableName() string {
	return "t_message_read"
}
