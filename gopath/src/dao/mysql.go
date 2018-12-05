package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/light4d/object4d/common/config"
)

func DB() *gorm.DB {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", config.APPConfig.Mysql)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}
	return db
}
