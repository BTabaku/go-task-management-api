package config

import (
	"go-task-management-api/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := AppConfig.DBSource // Ensure this is correctly formatted

	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Check if the database is reachable
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB object: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	// Migrate the schema
	if err := DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	log.Println("Connected to PostgreSQL and auto-migration completed.")
}
