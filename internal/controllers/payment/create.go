package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func MakePayment(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uuid.UUID)

	var payload models.CreatePaymentSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	payload.UserID = userId
	payment, err := services.CreatePayment(payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to create payment"})
	}

	//the actual checkout logic and moving book id to user's book list should be here but to make it easy to test it has been put in the checkout route

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": payment})
}

func GetPaymentByUserId(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	payment, err := services.GetPaymentsByUserID(userID.String())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get payment details"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": payment})
}

func GetPaymentById(c *fiber.Ctx) error {
	paymentId := c.Params("paymentId")

	payment, err := services.GetPaymentByID(paymentId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get payment details"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": payment})
}
