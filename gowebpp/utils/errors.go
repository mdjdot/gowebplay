package utils

import (
	"gowebpp/confs"
	"net/http"
)

// ProcessClieintError 处理错误的请求
func ProcessClientError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("错误的请求"))
	confs.Logger.Printf("请求信息不正确，请求：%v", r.Form)
}

// ProcessServerError 处理错误的响应
func ProcessServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("服务异常"))
	confs.Logger.Printf("服务异常，错误：%v", err)
}
