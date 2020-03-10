package dbs

import (
	"database/sql"

	"gowebp/logger"

	_ "github.com/go-sql-driver/mysql" // 初始化mysql
)

var (
	// WebDB mysql 数据库
	WebDB *sql.DB
	err   error
)

func init() {
	WebDB, err = sql.Open("mysql", "dm:dmtest@tcp(127.0.0.1:3306)/webdb")
	if err != nil {
		logger.Logger.Fatal(err)
	}

	err = WebDB.Ping()
	if err != nil {
		logger.Logger.Fatal(err)
	}
}
