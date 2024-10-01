// This file is responsible for loading and managing configuration settings for your application, such as database connection strings, server ports, and other environment variables.

package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port        string `json:"port"`
	DatabaseUrl string `json:"database_url"`
}

func LoadConfig(filePath string) (Config, error) {

	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)

	return config, err
}
