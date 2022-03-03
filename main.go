package main

import (
	"github.com/gin-gonic/gin"
	"mercury/app/api"
	"mercury/app/middleware"
	"mercury/app/subway"
)

func main() {
	subway.Start()

	route := gin.Default()
	route.Use(middleware.Cors())
	route.GET("/get_example", api.GetAccessToken)
	route.POST("/upload_canvas", api.UploadBinary)

	route.Run(":8085")
}
