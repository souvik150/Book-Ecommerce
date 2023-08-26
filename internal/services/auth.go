package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/config"
	database "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/utils"
)

func SignupUser(payload *models.RegisterUserSchema) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	otp, _ := utils.GenerateOTP(6)
	newUser := models.User{
		Username:     payload.Username,
		Email:        payload.Email,
		Password:     string(hashedPassword),
		PhoneNumber:  payload.PhoneNumber,
		ProfileImage: payload.ProfileImage,
		Verified:     false,
		Role:         payload.Role,
		Otp:          otp,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}

	body := fmt.Sprintf("Dear User,\n\nWelcome to the App! Thank you for joining us.\n\n"+
		"To complete your registration, please enter the following One-Time Password (OTP):\n\n"+
		"OTP: %s\n\n"+
		"This OTP is valid for a limited time only. Please keep it confidential and do not share it with anyone.\n\n", otp)

	msg := fmt.Sprintf("Books App\n%s", body)

	email, err := utils.SendEmail(payload.Email, msg)
	if err != nil {
		return err
	}
	fmt.Println(email)

	return nil
}

func LoginUser(payload *models.LoginUserSchema) (models.AuthResponse, error) {
	var user models.User
	result := database.DB.Where("email = ?", payload.Email).First(&user)
	if result.Error != nil {
		return models.AuthResponse{}, result.Error
	}

	fmt.Println(user)

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

func VerifyOTP(userID uuid.UUID, otp string) error {
	var user models.User
	result := database.DB.Where("id = ? AND otp = ?", userID, otp).First(&user)
	if result.Error != nil {
		return result.Error
	}

	// Update user's verification status and clear OTP
	user.Verified = true
	user.Otp = ""
	result = database.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GenerateAuthTokens(user *models.User) (models.AuthResponse, error) {
	access_token, err := utils.GenerateAccessToken(user)
	if err != nil {
		return models.AuthResponse{}, err
	}

	refresh_token, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return models.AuthResponse{}, err
	}

	// Store the refresh token in the database
	refreshTokenEntry := models.RefreshToken{
		UserID: user.ID,
		Token:  refresh_token,
	}
	result := database.DB.Create(&refreshTokenEntry)
	if result.Error != nil {
		return models.AuthResponse{}, result.Error
	}

	authResponse := models.AuthResponse{
		UserID:       user.ID,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		Verified:     user.Verified,
	}

	return authResponse, nil
}

func ResendOTP(userID uuid.UUID) error {
	otp, _ := utils.GenerateOTP(6)

	result := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(models.User{
		Otp: otp,
	})
	if result.Error != nil {
		return result.Error
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Verified == true {
		return nil
	}

	subject := "OTP Resent"
	body := fmt.Sprintf("Dear User,\n\nWe have resent the One-Time Password (OTP) to your email.\n\n"+
		"OTP: %s\n\n"+
		"This OTP is valid for a limited time only. Please keep it confidential and do not share it with anyone.\n\n"+
		"Thank you,\nThe Notes App Team", otp)

	msg := fmt.Sprintf("%s\n%s", subject, body)

	email, err := utils.SendEmail(user.Email, msg)
	if err != nil {
		return err
	}
	fmt.Println(email)

	return nil
}

func RefreshAccessToken(payload *models.RefreshTokenSchema) (models.AuthResponse, error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		return models.AuthResponse{}, err
	}

	// Check if the provided refresh token exists in the database
	var refreshTokenEntry models.RefreshToken
	result := database.DB.Where("token = ?", payload.RefreshToken).First(&refreshTokenEntry)
	if result.Error != nil {
		return models.AuthResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	// Parse and validate the access token
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(payload.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.RefreshTokenSecret), nil
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
		RefreshToken: payload.RefreshToken,
	}

	return authResponse, nil
}
