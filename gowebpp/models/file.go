package models

import (
	"github.com/astaxie/beego/orm"
)

// File 文件类型
type File struct {
	ID       int    `orm:"pk;auto;column(id)"`
	Name     string `orm:"size(30);column(name)"`
	Size     int64  `orm:"column(size)"`
	Hash     string `orm:"size(50);column(hash)"`
	Location string `orm:"size(255);column(location)"`
}

// Add 添加文件
func (f *File) Add() (int64, error) {
	// stmt, err := confs.DB.Prepare("insert into files (name, size, hash, location) values (?, ?, ?, ?)")
	// if err != nil {
	// 	return 0, err
	// }
	// defer stmt.Close()
	// results, err := stmt.Exec(f.Name, f.Size, f.Hash, f.Location)
	// if err != nil {
	// 	return 0, err
	// }
	// return results.LastInsertId()
	o := orm.NewOrm()
	return o.Insert(f)
}

// TableName 表名映射
func (f *File) TableName() string {
	return "files"
}
