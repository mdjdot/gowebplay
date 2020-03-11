package dbs

import (
	"fmt"
	"ginp/tables"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 注册mysql
	"github.com/spf13/viper"
)

// GinDB mysql数据库
var GinDB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)
	// db, err := gorm.Open("mysql", "dm:dmtest@tcp(127.0.0.1:3306)/gindb?charset=utf8&parseTime=true&loc=Local")
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic(err)
	}

	// db.CreateTable(&User{})
	db.AutoMigrate(&tables.User{})
	GinDB = db
}
