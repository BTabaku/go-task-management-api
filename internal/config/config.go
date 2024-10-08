package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBDriver string
	DBSource string
}

var AppConfig Config

// InitDatabase initializes the database connection
func InitDatabase() {
	dsn := AppConfig.DBSource
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to initialize database, got error %v\n", err)
		panic("failed to connect database")
	}
	DB = db
	fmt.Println("Database connection established")
}

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
	DB = database
}
