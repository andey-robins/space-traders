package repositories

import (
	"fmt"
	"sync"

	"github.com/andey-robins/space-traders/stback/models"
)

type ConfigRepository struct {
	db *SqliteDb
}

var configRepoInstance *ConfigRepository
var configRepoOnce sync.Once

func CreateConfigRepository() *ConfigRepository {
	configRepoOnce.Do(func() {
		configRepoInstance = &ConfigRepository{db: Connect()}
	})
	return configRepoInstance
}

func (cr *ConfigRepository) GetAll() ([]*models.Config, error) {
	q := "SELECT Id, Key, Value FROM Config"
	rows, err := cr.db.db.Query(q)
	check(err)
	defer rows.Close()

	configData := make([]*models.Config, 0)
	for rows.Next() {
		var cfg models.Config
		if err := rows.Scan(
			&cfg.Id,
			&cfg.Key,
			&cfg.Value,
		); err != nil {
			return nil, fmt.Errorf("error scanning config: %w", err)
		}

		configData = append(configData, &cfg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning configs: %w", err)
	}

	return configData, nil
}

func (cr *ConfigRepository) Save(config *models.Config) error {
	// updateQuery := "UPDATE Config SET Key = ?, Value = ? WHERE Id = ?"
	insertQuery := "INSERT INTO Config (Id, Key, Value) VALUES (?, ?, ?)"

	//_, err := cr.db.db.Exec(updateQuery, config.Key, config.Value, config.Id)
	_, err := cr.db.db.Exec(insertQuery, nil, config.Key, config.Value)

	return err
}
