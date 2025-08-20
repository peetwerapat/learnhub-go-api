package main

import (
	"github.com/gin-gonic/gin"
)

// @title My API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})

	r.Run()
}
