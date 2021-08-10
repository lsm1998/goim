package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"logic/model/message"
	"logic/model/user"
)

var dsn string

func initOrm() {
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		C.Mysql.User, C.Mysql.Password, C.Mysql.Host, C.Mysql.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		new(user.User),
		new(message.Message),
		new(message.Group),
		new(message.GroupAdmin))
	if err != nil {
		panic(err)
	}
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
