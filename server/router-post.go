package server

import (
	"log"
	"lowswoo/db"
	"lowswoo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// THIS
func createBucket(c *gin.Context) {

	bucket := models.BucketDB{}
	c.BindJSON(&bucket)

	err := db.CreateBucket(&bucket)
	if err != nil {
		c.JSON(http.StatusNotFound, bucket)
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, db.GetBucketListNames())

	}

}

func removeBucket(c *gin.Context) {
	bucket := struct {
		BucketName string `json:"bucketName"`
	}{}
	c.BindJSON(&bucket)
	err := db.RemoveBucket(bucket.BucketName)
	if err != nil {
		log.Default().Println(err)
	}
	c.JSON(http.StatusOK, db.GetBucketListNames())
}

func uploadFile(c *gin.Context) {
	file := models.FileUpload{}
	c.BindJSON(&file)
	links, err := db.UploadFile(&file, c.Request.Host)
	if err != nil {
		log.Default().Println(err)
	}
	c.JSON(http.StatusOK, links)
}

func setDescription(c *gin.Context) {
	d := models.BucketDescription{}
	c.BindJSON(&d)
	d.ID = "description"
	// log.Default().Println("Description ", d.Description)
	err := db.SetDescription(&d)
	if err != nil {
		log.Default().Println(err)
	}
	c.JSON(http.StatusOK, nil)
}
