package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Age  int    `json:age`
	Name string `json:name`
}

var people = []Person{
	{Age: 12, Name: "Da"},
	// {Age: 21, Name: "Fa"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func main() {
	router := gin.Default()
	router.GET("/people", getAlbums)

	router.Run("localhost:8080")
}
