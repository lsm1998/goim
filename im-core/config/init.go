package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var C *Config

type Config struct {
	Mysql
}

type Mysql struct {
	Url      string
	User     string
	Password string
	Db       string
}

func init() {
	C = new(Config)
	C.Mysql = Mysql{
		Url:      "119.29.117.244:3306",
		Db:       "im",
		User:     "root",
		Password: "fuck123",
	}
	// ---------Mysql--------
	db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, C.User, C.Password, C.Url, C.Db))
	if err != nil {
		panic(err)
	}
	DB = db
	DB.LogMode(true)
}
