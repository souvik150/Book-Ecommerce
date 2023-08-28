package review

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func GetReviewByUserID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	_, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	var reviews []models.Review
	results := database.DB.Find(&reviews, "user_id = ?", userID)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(reviews), "reviews": reviews})
}

func GetReviewsByBookID(c *fiber.Ctx) error {
	bookId := uuid.MustParse(c.Params("bookId"))

	var reviews []models.Review
	results := database.DB.Find(&reviews, "book_id = ?", bookId)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	var rating float64
	if len(reviews) > 0 {
		for _, review := range reviews {
			rating += review.Rating
		}
		rating = rating / float64(len(reviews))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(reviews), "reviews": reviews, "rating": rating})
}
