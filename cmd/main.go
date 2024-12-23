package main

import (
	"fmt"
	"virtual/configs"
	"virtual/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadFile("configs/configs.json")
}

func main() {
	r := gin.Default()

	uploadController := &controllers.UploadControllerResponse{}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.Use(gin.Logger())
	r.POST("/upload", uploadController.Upload)
	r.Static("/public", "./public")

	r.Run(fmt.Sprintf(":%d", configs.GetConfig().Server.Port))
}
