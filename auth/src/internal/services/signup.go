package services

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/models"
)

func SignupUser(payload *models.RegisterUserSchema) (models.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.AuthResponse{}, err
	}

	newUser := models.User{
		Username:     payload.Username,
		Email:        payload.Email,
		Password:     string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         "user",
		ProfileImage: payload.ProfileImage,
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return models.AuthResponse{}, result.Error
	}

	authResponse, err := GenerateAuthTokens(&newUser)
	if err != nil {
		return models.AuthResponse{}, err
	}

	return authResponse, nil
}
