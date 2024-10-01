package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver string
	DBSource string
}

var AppConfig Config

func LoadConfig(filePath string) (Config, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Println("No .env file found")
		return Config{}, err
	}

	config := Config{
		DBDriver: getEnv("DB_DRIVER", "sqlite"),
		DBSource: getEnv("DB_SOURCE", "tasks.db"),
	}

	AppConfig = config
	return config, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
