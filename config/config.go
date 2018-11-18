package config

import (
	"encoding/json"
	"os"
)

type DbConfig struct {
	DbType string `json:"DB_TYPE"`
	Name string `json:"DB_NAME"`
	Port int	`json:"DB_PORT"`
	User string `json:"DB_USER"`
	Pass string `json:"DB_PASS"`
}

func GetConfig(config interface{}) (error) {
	configFile, _ := os.Open("devConfig.json")
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err := decoder.Decode(config)
	if err != nil {
		return err
	}
	return err
}
