package services

import (
	"golang.org/x/crypto/bcrypt"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/models"
)

func LoginUser(payload *models.LoginUserSchema) (models.AuthResponse, error) {
	var user models.User
	result := database.DB.Where("username = ?", payload.Username).First(&user)
	if result.Error != nil {
		return models.AuthResponse{}, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return models.AuthResponse{}, err
	}

	authResponse, err := GenerateAuthTokens(&user)
	if err != nil {
		return models.AuthResponse{}, err
	}

	return authResponse, nil
}
