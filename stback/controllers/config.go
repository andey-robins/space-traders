package controllers

import (
	"net/http"
	"strconv"

	"github.com/andey-robins/space-traders/stback/models"
	"github.com/andey-robins/space-traders/stback/services"
	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	service *services.ConfigService
}

func (cc *ConfigController) CreateController(r *gin.Engine) *gin.Engine {
	configGroup := r.Group("/configs")
	{
		configGroup.GET("/all", cc.getAll)
		configGroup.PUT("/save", cc.save)
		configGroup.DELETE("/:configId", cc.delete)
	}

	cc.service = services.NewConfigService()

	return r
}

func (cc *ConfigController) getAll(c *gin.Context) {
	res, err := cc.service.GetAll()
	check(c, err)

	if err == nil {
		c.JSON(http.StatusOK, res)
	}
}

func (cc *ConfigController) save(c *gin.Context) {
	var config models.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.service.Save(&config)
	check(c, err)

	if err == nil {
		c.Status(http.StatusOK)
	}
}

func (cc *ConfigController) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("configId"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = cc.service.DeleteById(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
