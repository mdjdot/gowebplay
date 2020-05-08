package main

import (
	"net/http"

	"gowebpp/confs"
	"gowebpp/models"
	"gowebpp/routers"

	"github.com/astaxie/beego/orm"
)

func main() {
	// 初始化
	confs.InitLogger()
	confs.InitRedisPool()
	confs.InitMongo()
	confs.InitDB()
	orm.RegisterModel(&models.User{}, &models.File{})
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		confs.Logger.Fatalf("同步数据库失败，错误：%v\n", err)
	}

	// 创建多路复用器
	mux := http.NewServeMux()

	// 初始化路由
	routers.Init(mux)

	// 处理文件
	mux.Handle("/static/imgs", http.FileServer(http.Dir("./views/static/imgs")))

	// 启动服务并监听端口
	err = http.ListenAndServe(":8080", mux)
	confs.Logger.Println("程序启动：")
	if err != nil {
		confs.Logger.Fatalf("启动服务失败，错误：%v\n", err)
	}
}
