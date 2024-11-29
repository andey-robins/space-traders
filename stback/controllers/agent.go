package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentController struct{}

func (ac *AgentController) CreateController(r *gin.Engine) *gin.Engine {
	agentGroup := r.Group("/agents")
	{
		agentGroup.GET("/ping", agentPong)
	}

	return r
}

func agentPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
