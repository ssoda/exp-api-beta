package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", PingHandler)
	return r
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
