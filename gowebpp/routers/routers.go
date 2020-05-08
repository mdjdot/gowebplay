package routers

import (
	"gowebpp/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/websocket"
)

// Init 初始化路由
func Init(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.HandleHome)
	mux.HandleFunc("/login", controllers.HandleLogin)
	mux.HandleFunc("/signup", controllers.HandleSignup)
	mux.Handle("/echo", websocket.Handler(controllers.Echo))
}

// InitREST 初始化 Rest 路由
func InitREST(router *httprouter.Router) {
	router.GET("/:uid", controllers.RESTAssets)
}
