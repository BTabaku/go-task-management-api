package config

import (
	"go-task-management-api/internal/models"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDatabase() {
	once.Do(func() {
		var err error
		switch AppConfig.DBDriver {
		case "sqlite":
			db, err = gorm.Open(sqlite.Open(AppConfig.DBSource), &gorm.Config{})
		default:
			log.Fatalf("Unsupported DB driver: %s", AppConfig.DBDriver)
		}

		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		// Migrate the schema
		err = db.AutoMigrate(&models.Task{})
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
	})
}

func GetDB() *gorm.DB {
	return db
}
