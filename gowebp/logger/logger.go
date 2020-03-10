package logger

import (
	"log"
	"os"
)

var (
	// Logger 日志记录器
	Logger *log.Logger
)

func init() {
	_, err := os.Stat("./log/")
	if err != nil {
		err = os.Mkdir("./log/", os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile("./log/gowebp", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	Logger = log.New(file, "goweb: ", log.Ldate|log.Ltime|log.Llongfile)
}
