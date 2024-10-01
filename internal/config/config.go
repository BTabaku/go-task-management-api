package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	DBDriver string
	DBSource string
}

var (
	AppConfig Config
	db        *gorm.DB
)

// LoadConfig loads the configuration from the specified file
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

// getEnv gets the environment variable or returns the default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// SetDB sets the database connection
func SetDB(database *gorm.DB) {
	db = database
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
