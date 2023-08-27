package services

import (
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func GetCartById(cartId uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	result := database.DB.First(&cart, "id = ?", cartId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}

func GetCartByUserID(userId uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	result := database.DB.First(&cart, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}

func CreateCart(userId uuid.UUID) (*models.Cart, error) {
	cart := models.Cart{
		UserID:    userId,
		CreatedAt: database.DB.NowFunc(),
		UpdatedAt: database.DB.NowFunc(),
		Items:     []models.CartItem{},
	}
	result := database.DB.Create(&cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}
