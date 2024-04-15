package main

import "github.com/gin-gonic/gin"
import "net/http"


type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


var albums = []Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}



func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it's working fine !",
		})
	})

    router.Run("localhost:8080")
}