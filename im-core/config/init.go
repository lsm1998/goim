package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"im/model"
	"utils"
)

var DB *gorm.DB

var C *Config

type Config struct {
	Mysql
	Registry
	Rpc
}

type Mysql struct {
	Url      string
	User     string
	Password string
	Db       string
}

type Registry struct {
	Adders []string
}

type Rpc struct {
	Port     uint
	Server   string
	Metadata string
}

func init() {
	C = new(Config)
	err := utils.ScanConfig(C)
	if err != nil {
		panic(err)
	}

	// ---------Mysql--------
	db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, C.User, C.Password, C.Url, C.Db))
	if err != nil {
		panic(err)
	}
	DB = db
	DB.LogMode(true)

	utils.MigrateAuto(DB, []interface{}{
		&model.Message{},
	})
}
