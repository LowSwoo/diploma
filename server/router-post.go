package server

import (
	"log"
	minioServer "lowswoo/minio-server"
	"lowswoo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createBucket(c *gin.Context) {
	bucket := models.Bucket{}
	c.BindJSON(&bucket)
	minioServer.CreateBucket(&bucket)
	c.JSON(http.StatusOK, minioServer.GetBucketList())

}

func removeBucket(c *gin.Context) {
	bucket := struct {
		BucketName string `json:"bucketName"`
	}{}
	c.BindJSON(&bucket)
	err := minioServer.RemoveBucket(bucket.BucketName)
	if err != nil {
		log.Default().Println(err)
	}
	c.JSON(http.StatusOK, minioServer.GetBucketList())
}

func uploadFile(c *gin.Context) {
	type Link struct {
		FileName string `json:"fileName"`
		Url      string `json:"url"`
	}

	file := struct {
		BucketName string   `json:"bucketName"`
		FileName   []string `json:"fileName"`
	}{}
	c.BindJSON(&file)
	links := []Link{}
	for _, fn := range file.FileName {
		url := minioServer.GetFileLink(file.BucketName, fn)
		url_string := url.Scheme + "://" + url.Host + url.Path + "?" + url.RawQuery
		links = append(links, Link{FileName: fn, Url: url_string})
	}
	c.JSON(http.StatusOK, links)
}
