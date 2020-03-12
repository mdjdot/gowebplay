package main

import (
	"ginp/controls"
	"ginp/dbs"
	"ginp/middlewares"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	dbs.InitDB()

	r := gin.Default()
	{
		r.Use(middlewares.CROSSMiddleware())
		r.POST("/api/auth/register", controls.Register)
		r.GET("/api/auth/Login", controls.Login)
		r.GET("/api/auth/Info", middlewares.AuthMiddleware(), controls.Info)
	}
	port := viper.GetString("server.port")

	if port != "" {
		log.Fatal(r.Run(":" + port))
	}
	r.Run(":8080")
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}
