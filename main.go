package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mjbaig/file-manager/datastore"
)

func main() {

	awsProfile := os.Getenv("AWS_PROFILE")

	bucketName := os.Getenv("AWS_BUCKET")

	sess := datastore.CreateSessionInstace(awsProfile)

	uploader := datastore.CreateS3UploaderInstance(sess)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"hello": "test",
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		_, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.String(500, err.Error())
			return
		}

		log.Panicf("File: %s", fileHeader.Filename)

		file, err := fileHeader.Open()

		if err != nil {
			c.String(500, err.Error())
			return
		}

		datastore.UploadFile(file, fileHeader.Filename, bucketName, uploader)

		if err != nil {
			c.String(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"filename": fileHeader.Filename,
		})
	})

	r.Run(":8080")

}
