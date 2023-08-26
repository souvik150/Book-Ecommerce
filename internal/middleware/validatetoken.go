package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/config"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func validateAccessToken(c *fiber.Ctx) (*jwt.Token, error) {
	config, _ := config.LoadConfig(".")
	accessTokenSecretKey := config.AccessTokenSecret

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Authorization header missing")
	}

	accessToken := authHeader[len("Bearer "):]
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired access token")
	}

	return token, nil
}

func validateRole(c *fiber.Ctx, requiredRole string) (uuid.UUID, error) {
	token, err := validateAccessToken(c)
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}

	userIDStr, ok := claims["userID"].(string)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}

	var user models.User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not found")
	}

	if requiredRole != "" && user.Role != requiredRole {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not "+requiredRole)
	}

	return userID, nil
}

func UserTokenValidation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := validateRole(c, "")
		if err != nil {
			return err
		}
		return c.Next()
	}
}

func AdminTokenValidation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := validateRole(c, "admin")
		if err != nil {
			return err
		}
		return c.Next()
	}
}

func VendorTokenValidation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := validateRole(c, "vendor")
		if err != nil {
			return err
		}
		return c.Next()
	}
}
