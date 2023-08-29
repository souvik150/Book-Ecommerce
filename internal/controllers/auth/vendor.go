package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
)

func ApplyVendor(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	user, err := services.GetUserByID(uuid.MustParse(userID.String()))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	user.Role = "pending"
	user, err = services.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})

}

func ViewVendorApplications(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	user, err := services.GetUserByID(uuid.MustParse(userID.String()))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}
	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Only admins can view vendor applications"})
	}

	vendors, err := services.GetUsersByRole("pending")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant fetch applications"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"vendors": vendors}})
}

func ApproveVendor(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	applicantId := c.Params("userid")
	user, err := services.GetUserByID(uuid.MustParse(userID.String()))
	applicant, err := services.GetUserByID(uuid.MustParse(applicantId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Only admins can approve vendor applications"})
	}

	if applicant.Role != "pending" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "This user is not a pending vendor"})
	}

	applicant.Role = "vendor"
	applicant, err = services.UpdateUser(applicant)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": applicant}})
}

func RejectVendor(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	applicantId := c.Params("userid")
	user, err := services.GetUserByID(uuid.MustParse(userID.String()))
	applicant, err := services.GetUserByID(uuid.MustParse(applicantId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Only admins can approve vendor applications"})
	}

	if applicant.Role != "pending" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "This user is not a pending vendor"})
	}

	applicant.Role = "user"
	applicant, err = services.UpdateUser(applicant)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": applicant}})
}
