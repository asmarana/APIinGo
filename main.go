package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func AlbumbyID(c *gin.Context) {
	id := c.Param("id")
	albums, err := getAlbumbyId(id)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, albums)

}

func getAlbumbyId(id string) (*album, error) {
	for i, a := range albums {
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", AlbumbyID)
	router.Run("localhost:8080")
}
