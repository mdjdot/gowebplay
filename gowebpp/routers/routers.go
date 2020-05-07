package routers

import (
	"gowebpp/controllers"
	"net/http"
)

// Init 初始化路由
func Init(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.HandleHome)
	mux.HandleFunc("/login", controllers.HandleLogin)
	mux.HandleFunc("/signup", controllers.HandleSignup)
}
