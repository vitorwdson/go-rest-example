package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)

	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums/:id", updateAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid body format")
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.String(http.StatusNotFound, "Album not found")
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album

	newAlbum.ID = id
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid body format")
		return
	}

	for i, album := range albums {
		if album.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}
	}

	c.String(http.StatusNotFound, "Album not found")
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.String(http.StatusOK, "Album deleted successfully")
			return
		}
	}

	c.String(http.StatusNotFound, "Album not found")
}