package models

import "gowebpp/confs"

// File 文件类型
type File struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Size     int64  `json:"size,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Location string `json:"location,omitempty"`
}

// Add 添加文件
func (f *File) Add() (int64, error) {
	stmt, err := confs.DB.Prepare("insert into files (name, size, hash, location) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	results, err := stmt.Exec(f.Name, f.Size, f.Hash, f.Location)
	if err != nil {
		return 0, err
	}
	return results.LastInsertId()
}
