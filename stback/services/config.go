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
	oldConfig, err := cs.repo.GetById(config.Id)
	if err != nil {
		return err
	}

	if oldConfig == nil {
		return cs.repo.Create(config)
	} else {
		return cs.repo.Update(config)
	}
}

func (cs *ConfigService) Delete(config *models.Config) error {
	return cs.repo.Delete(config.Id)
}

func (cs *ConfigService) DeleteById(id models.ConfigId) error {
	return cs.repo.Delete(id)
}
