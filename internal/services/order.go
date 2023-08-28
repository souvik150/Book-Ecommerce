package services

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func GetOrders(orderIDs pq.StringArray) ([]models.Order, error) {
	var orders []models.Order
	for _, orderID := range orderIDs {
		var order models.Order
		err := database.DB.First(&order, "id = ?", orderID).Error
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderByID(orderID string) (*models.Order, error) {
	orderuid := uuid.MustParse(orderID)
	var order models.Order
	err := database.DB.First(&order, "id = ?", orderuid).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrderByCartID(cartID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Where("cart_id = ?", cartID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
