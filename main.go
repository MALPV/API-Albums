package main

import (
	"malpvapps/web-services-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", routers.GetAlbums)
	router.GET("/albums/:id", routers.GetAlbumByID)
	router.POST("/albums", routers.PostAlbums)

	router.Run("localhost:8080")
}
