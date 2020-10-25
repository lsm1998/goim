package model

import "utils"

type User struct {
	utils.Model
	Nickname string `json:"nickname" gorm:"nickname"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Salt     string `json:"salt" gorm:"salt"`
	HeadImg  string `json:"head_img" gorm:"head_img"`
}

func (User) TableName() string {
	return "t_user"
}
