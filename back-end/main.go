package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	router.Run(":8000")
}
