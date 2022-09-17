package server

import (
	minioServer "lowswoo/minio-server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Return list of buckets as a list
func getBucketList(c *gin.Context) {
	c.JSON(http.StatusOK, minioServer.GetBucketList())
}

func getFileList(c *gin.Context) {
	c.JSON(http.StatusOK, minioServer.GetFileList(c.Query("bucketName")))
}

func removeFile(c *gin.Context) {
	minioServer.RemoveFile(c.Query("bucketName"), c.Query("fileName"))
	c.JSON(http.StatusOK, minioServer.GetFileList(c.Query("bucketName")))
}
