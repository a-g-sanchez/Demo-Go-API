package controllers

import (
	"github.com/a-g-sanchez/go-crud/models"
	"github.com/gin-gonic/gin"
)

// GET all the albums in the database
func GetAlbums(c *gin.Context) {
	var albums []models.Album

	models.DB.Find(&albums)

	c.IndentedJSON(200, gin.H{"data": albums})
}

// GET one album off the id
func FindAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.IndentedJSON(404, gin.H{"error": "Album instance not found!"})
	}

	c.IndentedJSON(200, gin.H{"data": album})
}

// Create a struct that will look to make sure the req.body is correct
type AddAlbumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

// POST an album to the database
func AddAlbum(c *gin.Context) {
	var input AddAlbumInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}

	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	models.DB.Create(&album)

	c.IndentedJSON(200, gin.H{"data": album})
}

// Create a struct to control the shape of req.body
type UpdateAlbumInput struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// PUT request to update an album
func UpdateAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id=?", c.Param("id")).First(&album).Error; err != nil {
		c.IndentedJSON(404, gin.H{"error": "Album instance not found"})
		return
	}

	var input UpdateAlbumInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(404, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(input)

	c.IndentedJSON(200, gin.H{"data": album})

}

func DeleteAlbum(c *gin.Context) {

	var album models.Album

	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.IndentedJSON(404, gin.H{"error": "Album record not found"})
		return
	}

	models.DB.Delete(&album)

	c.IndentedJSON(200, gin.H{"data": true})
}
