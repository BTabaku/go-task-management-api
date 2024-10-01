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
	dsn := AppConfig.DBSource
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	DB.AutoMigrate(&models.Task{})
}
