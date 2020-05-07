package confs

import (
	"log"
	"os"
)

// Logger 记录器
var Logger *log.Logger

// InitLogger 初始化记录器
func InitLogger() {
	f, err := os.OpenFile("hello.log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	Logger = log.New(f, "", log.LstdFlags)
}
