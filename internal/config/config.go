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
		DBDriver: getEnv("DB_DRIVER", "postgres"),
		DBSource: getEnv("DB_SOURCE", "host=localhost user=task_manager password=SecureP@ssw0rd! dbname=task_management port=5432 sslmode=disable"),
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
