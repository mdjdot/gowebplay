package routers

import (
	"ginp/controllers"
	"ginp/middlewares"

	"github.com/gin-gonic/gin"
)

// InitRouters 初始化路由
func InitRouters(r *gin.Engine) {
	r.POST("/api/auth/register", controllers.Register)
	r.GET("/api/auth/Login", controllers.Login)
	r.GET("/api/auth/Info", middlewares.AuthMiddleware(), controllers.Info)
}
