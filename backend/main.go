package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/private", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "private"})
	})

	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "admin"})
	})

	r.Run()
}
