package review

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func DeleteReview(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	reviewID := uuid.MustParse(c.Params("reviewId"))

	// find review
	review, err := services.GetReviewByID(reviewID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get review details"})
	}

	// check if user is the owner of the review
	if review.UserID != userID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not authorized to delete this review"})
	}

	// delete review
	err = services.DeleteReview(reviewID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to delete review"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Review deleted successfully"})

}
