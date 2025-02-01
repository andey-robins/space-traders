package repositories

import (
	"database/sql"
	"fmt"
	"log"
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

func (cr *ConfigRepository) Create(config *models.Config) error {
	log.Println("create config")
	q := "INSERT INTO Config (Key, Value) VALUES (?, ?)"
	_, err := cr.db.db.Exec(q, config.Key, config.Value)
	return err
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

func (cr *ConfigRepository) GetById(id models.ConfigId) (*models.Config, error) {
	q := "SELECT Id, Key, Value FROM Config WHERE Id = ?"
	row := cr.db.db.QueryRow(q, id)

	var cfg models.Config
	if err := row.Scan(
		&cfg.Id,
		&cfg.Key,
		&cfg.Value,
	); err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("error scanning config: %w", err)
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return &cfg, nil
}

func (cr *ConfigRepository) Update(config *models.Config) error {
	log.Println("update config")
	q := "UPDATE Config SET Key = ?, Value = ? WHERE Id = ?"
	_, err := cr.db.db.Exec(q, config.Key, config.Value, config.Id)
	return err
}

func (cr *ConfigRepository) Delete(id models.ConfigId) error {
	log.Println("deleting config")
	q := "DELETE FROM Config WHERE Id = ?"
	_, err := cr.db.db.Exec(q, id)
	return err
}
