package message

import "gorm.io/gorm"

type Group struct {
	gorm.Model
}

func (*Group) TableName() string {
	return "t_group"
}

type GroupAdmin struct {
	gorm.Model
}

func (*GroupAdmin) TableName() string {
	return "t_group_admin"
}
