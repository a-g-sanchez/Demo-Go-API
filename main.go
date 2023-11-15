package main

import (
	"github.com/a-g-sanchez/go-crud/controllers"
	"github.com/a-g-sanchez/go-crud/initializers"
	"github.com/a-g-sanchez/go-crud/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
}

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.FindAlbum)

	router.POST("/albums", controllers.AddAlbum)

	router.PUT("/albums/:id", controllers.UpdateAlbum)

	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	router.Run()
}
