package stback

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Get("/ping", pong)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
