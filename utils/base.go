package utils

import "time"

type Model struct {
	Id         int64      `json:"id" gorm:"id"`
	UpdateBy   int64      `json:"update_by" gorm:"update_by"`
	CreateBy   int64      `json:"create_by" gorm:"create_by"`
	UpdateTime time.Time  `json:"update_time" gorm:"update_time"`
	CreateTime time.Time  `json:"create_time" gorm:"create_time"`
	DeleteTime *time.Time `json:"delete_time" gorm:"delete_time"`
}

type PageInfo struct {
	Page uint32 `json:"page"`
	Size uint32 `json:"size"`
}

func (p *PageInfo) Offset() uint32 {
	return (p.Page - 1) * p.Size
}
