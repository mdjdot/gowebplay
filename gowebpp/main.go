package main

import (
	"net/http"

	"gowebpp/confs"
	"gowebpp/routers"
)

func main() {
	// 初始化
	confs.InitLogger()
	confs.InitDB()

	// 创建多路复用器
	mux := http.NewServeMux()

	// 初始化路由
	routers.Init(mux)

	// 处理文件
	mux.Handle("/static/imgs", http.FileServer(http.Dir("./views/static/imgs")))

	// 启动服务并监听端口
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		confs.Logger.Fatalf("启动服务失败，错误：%v", err)
	}
}
