package server

import (
	"log"
	"lowswoo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Return list of buckets as a list
func getBucketList(c *gin.Context) {
	c.JSON(http.StatusOK, db.GetBucketListNames())
}

// Returns list of files on the bucket as a list
func getFileList(c *gin.Context) {
	c.JSON(http.StatusOK, db.GetFileList(c.Query("bucketName")))
}

// Remove file from minio server
func removeFile(c *gin.Context) {
	db.RemoveFile(c.Query("bucketName"), c.Query("fileName"))
	log.Default().Println(c.Query("fileName"))
	c.JSON(http.StatusOK, db.GetFileList(c.Query("bucketName")))
}

func downloadFile(c *gin.Context) {
	addr := db.DownloadFile(c.Query("bucketName"), c.Query("filename"))
	log.Default().Println(addr)
	c.JSON(http.StatusOK, addr)
}
