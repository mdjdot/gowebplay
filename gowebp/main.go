package main

import (
	"gowebp/controls"
	"gowebp/logger"
	"net/http"
)

func main() {
	logger.Logger.Println("goweb app start")

	controls.InitControls()

	logger.Logger.Fatal(http.ListenAndServe(":8080", nil))
}
