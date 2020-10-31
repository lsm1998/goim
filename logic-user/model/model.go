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
