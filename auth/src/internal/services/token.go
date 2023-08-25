package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/utils"
)

func GenerateAuthTokens(user *models.User) (models.AuthResponse, error) {
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		return models.AuthResponse{}, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return models.AuthResponse{}, err
	}

	// Store the refresh token in the database
	refreshTokenEntry := models.RefreshToken{
		UserID: user.ID,
		Token:  refreshToken,
	}
	result := database.DB.Create(&refreshTokenEntry)
	if result.Error != nil {
		return models.AuthResponse{}, result.Error
	}

	authResponse := models.AuthResponse{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Verified:     true,
	}

	return authResponse, nil
}

func RefreshAccessToken(refreshToken string) (models.AuthResponse, error) {
	refrshSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	// Check if the provided refresh token exists in the database
	var refreshTokenEntry models.RefreshToken
	result := database.DB.Where("token = ?", refreshToken).First(&refreshTokenEntry)
	if result.Error != nil {
		return models.AuthResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	// Parse and validate the access token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(refrshSecret), nil
	})
	if err != nil {
		return models.AuthResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	// Generate a new access token
	userIDStr, ok := claims["userID"].(string)
	if !ok {
		return models.AuthResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return models.AuthResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	var user models.User
	user.ID = userID
	accessToken, err := utils.GenerateAccessToken(&user)
	if err != nil {
		return models.AuthResponse{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate access token")
	}

	authResponse := models.AuthResponse{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return authResponse, nil
}
