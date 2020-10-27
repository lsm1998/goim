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
