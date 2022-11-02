package server

import (
	"log"
	server "lowswoo/minio-server"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router *gin.Engine
	host   string
	port   string
)

func Run() {
	initServer()
	initializeRoutes()
	router.Run(string(host + ":" + port))
}

func initServer() {
	router = gin.Default()
	setServerSettings()
	router.Use(static.Serve("/", static.LocalFile("front/dist", true)))
	router.NoRoute(func(ctx *gin.Context) {
		if !strings.HasPrefix(ctx.Request.RequestURI, "/api") {
			ctx.File("/front/dist/index.html")
		}
	})
}

func setServerSettings() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
	host = os.Getenv("SERVER_HOST")
	port = os.Getenv("SERVER_PORT")
}

func initializeRoutes() {

	router.Use(authMiddleware)
	api := router.Group("/api")
	api.Use(SetHost)
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
		// file.GET("/download", downloadFile)
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

func SetHost(c *gin.Context) {
	server.SetHost(c.Request.Host)
}
