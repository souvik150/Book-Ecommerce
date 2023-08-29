package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running Migrations")

	err := db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Book{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.Review{},
		&models.Payment{},
	)
	if err != nil {
		fmt.Println("Migration error")
		return
	}

	log.Println("🚀 Migrations completed")
}
