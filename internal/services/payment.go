package services

import (
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func CreatePayment(payload models.CreatePaymentSchema) (models.Payment, error) {
	order := models.Payment{
		UserID:        payload.UserID,
		OrderID:       payload.OrderID,
		PaymentID:     payload.PaymentID,
		Signature:     payload.Signature,
		PaymentStatus: payload.PaymentStatus,
	}

	err := database.DB.Create(&order).Error
	if err != nil {
		return models.Payment{}, err
	}

	return order, nil
}

func GetPaymentsByUserID(userID string) ([]models.Payment, error) {
	var payments []models.Payment
	result := database.DB.Find(&payments, "user_id = ?", userID)
	if result.Error != nil {
		return []models.Payment{}, result.Error
	}

	return payments, nil
}

func GetPaymentByID(paymentID string) (models.Payment, error) {
	var payment models.Payment
	result := database.DB.First(&payment, "id = ?", paymentID)
	if result.Error != nil {
		return models.Payment{}, result.Error
	}

	return payment, nil
}
