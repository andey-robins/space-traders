package services

import (
	"sync"

	"github.com/andey-robins/space-traders/stback/models"
	"github.com/andey-robins/space-traders/stback/repositories"
)

type ConfigService struct {
	repo *repositories.ConfigRepository
}

var configServiceInstance *ConfigService
var configServiceOnce sync.Once

func NewConfigService() *ConfigService {
	configServiceOnce.Do(func() {
		configServiceInstance = &ConfigService{repo: repositories.CreateConfigRepository()}
	})
	return configServiceInstance
}

func (cs *ConfigService) GetAll() ([]*models.Config, error) {
	return cs.repo.GetAll()
}

func (cs *ConfigService) Save(config *models.Config) error {
	return cs.repo.Save(config)
}
