package utils

import "github.com/jinzhu/gorm"

func MigrateAuto(db *gorm.DB, tables []interface{}) {
	logDB := db.Debug()
	var err error
	for _, v := range tables {
		if !db.HasTable(v) {
			err = logDB.CreateTable(v).Error
		} else {
			err = logDB.AutoMigrate(v).Error
		}
		if err != nil {
			panic(err)
		}
	}
}
