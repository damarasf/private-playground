package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	})
}

var MovieGetHandler = func(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if movie, ok := movies[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"data": movie,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "movie not found",
		})
	}
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	movies := router.Group("/movie")
	{
		movies.GET("/list", MovieListHandler)
		movies.GET("/get/:id", MovieGetHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
