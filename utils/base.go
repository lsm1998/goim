package utils

import "time"

type Model struct {
	Id         int64      `json:"id" gorm:"id"`
	UpdateTime time.Time `json:"update_time" gorm:"update_time"`
	CreateTime time.Time `json:"create_time" gorm:"create_time"`
	DeleteTime *time.Time `json:"delete_time" gorm:"delete_time"`
}
