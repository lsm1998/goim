package model

import "utils"

type Role struct {
	utils.Model
	RoleName string `json:"role_name" gorm:"role_name"`
}

func (Role) TableName() string {
	return "t_role"
}

type Permission struct {
	utils.Model
	RoleId int64  `json:"role_id" gorm:"role_id"`
	Path   string `json:"path" gorm:"path"`
}

func (Permission) TableName() string {
	return "t_permission"
}

type Auth struct {
	utils.Model
	UserId int64 `json:"user_id" gorm:"user_id"`
	RoleId int64 `json:"role_id" gorm:"role_id"`
}

func (Auth) TableName() string {
	return "t_auth"
}

type User struct {
	utils.Model
	Nickname string `json:"nickname" gorm:"nickname"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Salt     string `json:"salt" gorm:"salt"`
	HeadImg  string `json:"head_img" gorm:"head_img"`
	AesKey   string `json:"aes_key" gorm:"aes_key"`
}

func (User) TableName() string {
	return "t_user"
}

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
	GroupId int64 `json:"group_id" gorm:"group_id"`
	UserId  int64 `json:"user_id" gorm:"user_id"`
}

func (GroupItem) TableName() string {
	return "t_group_item"
}

type Friends struct {
	utils.Model
	GroupId   int64 `json:"group_id" gorm:"group_id"`
	UserId    int64 `json:"user_id" gorm:"user_id"`
	FriendsId int64 `json:"friends_id" gorm:"friends_id"`
}

func (Friends) TableName() string {
	return "t_friends"
}
