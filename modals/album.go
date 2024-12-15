package modals

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// album represents data about a record album.
// @Description Album data representation
type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// var albums = []Album{
//     {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// // GetAlbums godoc
// // @Summary Get a list of albums
// // @Description Get all albums available in the store
// // @Tags albums
// // @Accept json
// // @Produce json
// // @Success 200 {array} Album
// // @Router /albums [get]
// func GetAlbums(c *gin.Context) {
//     c.JSON(http.StatusOK, albums)
// }
