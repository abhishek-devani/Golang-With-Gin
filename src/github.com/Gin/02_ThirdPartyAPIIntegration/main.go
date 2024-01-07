package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/get", getValues)

	r.Run(":9090")

}

var url = "http://date.jsontest.com"

func getValues(c *gin.Context) {

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	var result map[string]interface{}

	err = json.Unmarshal(data, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
		"date": result["date"],
	})

}