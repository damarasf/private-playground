package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "data not found",
			})
			return
		}
		user := data[idInt]
		c.JSON(http.StatusOK, gin.H{
			"name":    user.Name,
			"country": user.Country,
			"age":     user.Age,
		})
	}
}

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/profile/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusNotFound, "data not found")
			return
		}
		user := data[id]
		if user.Name == "" {
			c.String(http.StatusNotFound, "data not found")
			return
		}
		c.String(http.StatusOK, "Name: %s, Country: %s, Age: %d", user.Name, user.Country, user.Age)
		fmt.Printf("Name: %s, Country: %s, Age: %d", user.Name, user.Country, user.Age)
	})
	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
