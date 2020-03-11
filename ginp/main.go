package main

import (
	"ginp/controls"
	"ginp/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	{
		r.POST("/api/auth/register", controls.Register)
		r.GET("/api/auth/Login", controls.Login)
		r.GET("/api/auth/Info", middlewares.AuthMiddleware(), controls.Info)
	}
	r.Run(":8080")
}
