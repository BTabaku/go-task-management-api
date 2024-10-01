package config

import (
	"testing"
)

func TestInitDatabase(t *testing.T) {
	AppConfig = Config{
		DBDriver: "postgres",
		DBSource: "host=localhost user=your_user password=your_password dbname=tasks_db port=5432 sslmode=disable",
	}
	InitDatabase()
	if DB == nil {
		t.Fatal("Expected DB to be initialized, got nil")
	}
	// Verify the connection
	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer sqlDB.Close()
	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
