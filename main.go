package main

import (
	"api-project/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "api-project/docs"
)

func main () {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.AddAlbum)

	router.Run(":8080")
}