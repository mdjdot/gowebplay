package main

import (
	"fmt"
	"gowebpp/confs"
	"gowebpp/models"
	"gowebpp/routers"
	m "gowebpp/tools/models"
	"net/http"
	"net/rpc/jsonrpc"

	"github.com/astaxie/beego/orm"
	"github.com/julienschmidt/httprouter"
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

	go func() {
		// 使用httprouter
		router := httprouter.New()
		routers.InitREST(router)
		err = http.ListenAndServe(":8081", router)
		if err != nil {
			confs.Logger.Fatalf("启动服务失败，错误：%v\n", err)
		}
	}()

	{
		// client, err := rpc.DialHTTP("tcp", "127.0.0.1:8082")
		// client, err := rpc.Dial("tcp", "127.0.0.1:8082")
		client, err := jsonrpc.Dial("tcp", "127.0.0.1:8082")
		if err != nil {
			confs.Logger.Fatalf("rpc服务调用失败，错误：%v\n", err)
		}
		req := &m.Request{A: 10, B: 20}
		resp := &m.Response{}
		client.Call("Arith.Sum", req, resp)
		fmt.Println(resp.C)
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
