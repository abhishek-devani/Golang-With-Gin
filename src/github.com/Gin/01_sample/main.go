package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type api struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var data api

func main() {
	
	// Create router
	r := gin.Default()

	// HTTP Methods
	r.GET("/get", getValues)
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)

	// Start the engine
	r.Run(":9090")

}

func getValues(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func postValues(c *gin.Context) {

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something Wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func putValues(c *gin.Context) {

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something Wrong",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteValues(c *gin.Context) {

	data = api{}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}