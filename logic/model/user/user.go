package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
}

func (*User) TableName() string {
	return "t_user"
}
