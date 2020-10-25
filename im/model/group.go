package model

import "utils"

type Group struct {
	utils.Model
	// Leader 群主
	Leader int64 `json:"leader" gorm:"leader"`
	// GroupName 群名称
	GroupName string `json:"group_name" gorm:"group_name"`
	// GroupImg 群名称
	GroupImg string `json:"group_img" gorm:"group_img"`
}

func (Group) TableName() string {
	return "t_group"
}

type GroupItem struct {
	utils.Model
	GroupId int64
	UserId  int64
}

func (GroupItem) TableName() string {
	return "t_group_item"
}
