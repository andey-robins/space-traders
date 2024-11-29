package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(c *gin.Context, err error) {
	if err != nil {
		c.String(http.StatusInternalServerError, "error: %s", err.Error())
	}
}

type Controller interface {
	CreateController(r *gin.Engine) *gin.Engine
}

func RegisterControllers(r *gin.Engine) *gin.Engine {
	controllers := []Controller{
		&AgentController{},
		&ConfigController{},
	}

	for _, controller := range controllers {
		r = controller.CreateController(r)
	}

	return r
}
