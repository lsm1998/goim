package message

import "gorm.io/gorm"

type Message struct {
	gorm.Model
}

func (*Message) TableName() string {
	return "t_message"
}
