package dbs

import (
	"ginp/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GinDB mysql数据库
var GinDB *gorm.DB

// InitDB 初始化数据库
func init() {
	db, err := gorm.Open("mysql", "dm:dmtest@tcp(127.0.0.1:3306)/gindb?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	// db.CreateTable(&User{})
	db.AutoMigrate(&models.User{})
	GinDB = db
}
