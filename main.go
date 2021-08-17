package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mjbaig/file-manager/datastore"
)

func main() {

	datastore.UploadFile("./properties.yml", "hiimmaz-file-manager", "properties.yml")

	log.Println("Hello, World!")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "test",
		})
	})

	r.Run(":8080")

}
