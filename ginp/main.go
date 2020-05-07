package main

import (
	"ginp/dbs"
	"ginp/middlewares"
	"ginp/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	dbs.InitDB()

	r := gin.Default()
	r.Use(middlewares.CROSSMiddleware())
	routers.InitRouters(r)

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
