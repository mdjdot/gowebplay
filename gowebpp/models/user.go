package models

import (
	"github.com/astaxie/beego/orm"
)

// User 用户类型
type User struct {
	ID       int    `orm:"pk;auto;column(id)"`
	Name     string `orm:"size(30);column(name)"`
	Password string `orm:"size(50);column(password)"`
}

// IsExist 用户名是否存在
func (u *User) IsExist() bool {
	// stmt, err := confs.DB.Prepare("select name from users where name=?")
	// if err != nil {
	// 	confs.Logger.Printf("准备查询错误：%v\n", err)
	// }
	// defer stmt.Close()
	// results, err := stmt.Query(u.Name)
	// if err != nil {
	// 	confs.Logger.Printf("执行查询错误：%v\n", err)
	// }
	// defer results.Close()
	// if results.Next() {
	// 	return true
	// }
	// return false
	o := orm.NewOrm()
	o.Read(u, "name")
	if u.ID > 0 {
		return true
	}
	return false
}

// Add 添加用户
func (u *User) Add() (int64, error) {
	// stmt, err := confs.DB.Prepare("insert into users (name, password) values(?, ?)")
	// if err != nil {
	// 	confs.Logger.Printf("准备查询错误：%v\n", err)
	// 	return 0, err
	// }
	// defer stmt.Close()
	// results, err := stmt.Exec(u.Name, u.Password)
	// if err != nil {
	// 	confs.Logger.Printf("执行查询错误：%v\n", err)
	// 	return 0, err
	// }
	// return results.LastInsertId()
	o := orm.NewOrm()
	return o.Insert(u)
}

// GetByNameAndPwd 获取用户
func (u *User) GetByNameAndPwd() {
	// stmt, err := confs.DB.Prepare("select id from users where name=? and password=?")
	// if err != nil {
	// 	confs.Logger.Printf("准备查询错误：%v\n", err)
	// }
	// defer stmt.Close()
	// results, err := stmt.Query(u.Name, u.Password)
	// if err != nil {
	// 	confs.Logger.Printf("执行查询错误：%v\n", err)
	// }
	// defer results.Close()
	// if results.Next() {
	// 	results.Scan(&u.ID)
	// }
	o := orm.NewOrm()
	o.Read(u, "name", "password")

}

// TableName 表名映射
func (u *User) TableName() string {
	return "users"
}
