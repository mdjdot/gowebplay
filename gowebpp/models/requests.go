package models

import (
	"gowebpp/confs"
	"net/url"
)

// Request 请求类型
type Request struct {
	URL  *url.URL `json:"url,omitempty"`
	Time int64    `json:"time,omitempty"`
}

// Insert 插入数据
func (r *Request) Insert() error {
	col := confs.MongoConn.DB("appdb").C("logs")
	err := col.Insert(r)
	return err
}
