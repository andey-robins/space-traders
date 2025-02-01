package models

type ConfigId = int

type Config struct {
	Id    ConfigId `json:"id"`
	Key   string   `json:"key"`
	Value string   `json:"value"`
}
