package review

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	reviewSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/review"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func CreateReview(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	bookID := uuid.MustParse(c.Params("bookId"))

	// Get payload
	var payload reviewSchema.CreateReviewSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// find user
	user, err := services.GetUserByID(userID)

	if payload.Rating < 1 || payload.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Rating must be greater than 0 and less than 5"})
	}

	if payload.Comment == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Review cannot be empty"})
	}

	var review models.Review
	result, err := services.GetReviewByUserIDAndBookID(userID, bookID)
	if err != nil {
		// Handle other errors, if any, but no need to handle "record not found"
		if err != gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get review details"})
		}
	}

	// If review exists, update it; otherwise, create a new review
	if result.ID != uuid.Nil {
		review, err = services.UpdateReview(userID, bookID, user.Username, payload)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to update review"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": review})
	}

	// If no review exists, create a new review
	review, err = services.CreateReview(userID, bookID, user.Username, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to create review"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": review})

}
