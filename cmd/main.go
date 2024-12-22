package main

import (
	"fmt"
	"virtual/configs"
	"virtual/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadFile("configs/configs.json")
}

func main() {
	r := gin.Default()

	uploadController := &controllers.UploadControllerResponse{}
	r.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
	r.Use(gin.Logger())
	r.POST("/upload", uploadController.Upload)
	r.Static("/public", "./public")

	r.Run(fmt.Sprintf(":%d", configs.GetConfig().Server.Port))
}
