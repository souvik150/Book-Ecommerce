package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
)

func TokenValidation(c *fiber.Ctx) error {
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Authorization header missing"})
	}

	accessToken := authHeader[len("Bearer "):]
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid or expired access token"})
	}

	userIDStr, ok := claims["userID"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid access token"})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid access token"})
	}

	// Pass the authenticated user ID and user role to the context
	c.Locals("userID", userID)

	return c.Next()
}
