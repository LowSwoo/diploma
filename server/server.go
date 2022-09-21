package server

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Run() {
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("frontend", false)))
	initializeRoutes()
	router.Run("localhost:8080")
}

func initializeRoutes() {

	router.Use(authMiddleware)
	api := router.Group("/api")
	bucket := api.Group("/bucket")
	{
		bucket.GET("/list", getBucketList)
		bucket.POST("/remove", removeBucket)
		bucket.POST("/create", createBucket)

	}

	file := api.Group("/file")
	{
		file.POST("/upload", uploadFile)
		file.GET("/list", getFileList)
		file.GET("/remove", removeFile)
	}
}

// Middleware is temporary solution. No access control needed when fronted served by golang server
func authMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
