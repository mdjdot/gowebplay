package main

import (
	"ginp/controls"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", controls.Register)
	r.Run(":8080")
}
