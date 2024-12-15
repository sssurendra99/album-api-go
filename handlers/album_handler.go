package handlers

import (
	"api-project/db"
	"api-project/modals"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAlbums godoc
// @Summary Get a list of albums
// @Description Get all albums available in the store
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} modals.Album
// @Router /albums [get]
func GetAlbums(c *gin.Context) {

	db.CreateAlbumTable()
	// Fetch albums from the database
	albums, err := db.GetAlbums()
	if err != nil {
		log.Printf("Error fetching albums: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve albums"})
		return
	}

	// Return the albums in JSON format
	c.JSON(http.StatusOK, albums)
}

// AddAlbum godoc
// @Summary Add a new album
// @Description Add a new album to the store
// @Tags albums
// @Accept json
// @Produce json
// @Param album body modals.Album true "Album data"
// @Success 201 {object} modals.Album
// @Router /albums [post]
func AddAlbum(c *gin.Context) {
	var album modals.Album

	// Bind the JSON body to the album struct
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Insert the album into the database
	err := db.InsertAlbum(album.Title, album.Artist, album.Price)
	if err != nil {
		log.Printf("Error inserting album: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add album"})
		return
	}

	// Return the created album
	c.JSON(http.StatusCreated, album)
}