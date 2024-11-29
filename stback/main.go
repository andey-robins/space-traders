package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", pong)

	router.Run(":8080")
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
