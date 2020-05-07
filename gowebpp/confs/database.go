package confs

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // 注册 sqlite 数据库驱动
)

// // DB 数据库
// var DB *sql.DB

// InitDB 初始化数据库
func InitDB() {
	// db, err := sql.Open("sqlite3", "./datas/data.db")
	// if err != nil {
	// 	Logger.Fatalf("初始化数据库失败，错误：%v", err)
	// }
	// db.SetConnMaxLifetime(10 * time.Second)
	// db.SetMaxOpenConns(10)
	// DB = db

	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./datas/newdata.db")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.Debug = true
}
