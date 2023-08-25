package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=auth-postgres-srv user=admin password=password123 dbname=golang_fiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database!\n", err.Error())
		os.Exit(1)
	}

	// Enable UUID extension
	result := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if result.Error != nil {
		log.Fatal("Failed to enable UUID extension!\n", result.Error)
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	if err := DB.AutoMigrate(&models.User{}, &models.RefreshToken{}); err != nil {
		log.Fatal("Failed to run migrations!\n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")
}
