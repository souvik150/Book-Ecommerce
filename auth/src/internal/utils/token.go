package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/auth/src/internal/models"

	"time"

	"github.com/google/uuid"
)

func generateToken(userID uuid.UUID, secretKey string, expiration time.Duration, user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID.String(),
		"exp":       time.Now().Add(expiration).Unix(),
		"username":  user.Username,
		"role":      user.Role,
		"email":     user.Email,
		"createdAt": user.CreatedAt.Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func GenerateAccessToken(user *models.User) (string, error) {
	accessExpiry, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	return generateToken(user.ID, accessSecret, time.Duration(accessExpiry), user)
}

func GenerateRefreshToken(user *models.User) (string, error) {
	refreshExpiry, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY"))
	refrshSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	return generateToken(user.ID, refrshSecret, time.Duration(refreshExpiry), user)
}
