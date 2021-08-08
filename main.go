package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Hello, World!")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.Run(":8080")

}
