package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mime/multipart"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/utils"
)

func SignupUser(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Failed to process form data"})
	}

	// Get user data from the form
	username := form.Value["username"][0]
	email := form.Value["email"][0]
	password := form.Value["password"][0]
	phoneNo := form.Value["phone_number"][0]
	role := "user"

	if username == "" || email == "" || password == "" || phoneNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request body"})
	}

	_, err = services.GetUserByEmail(email)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "User already exists"})
	}

	// Handle profile picture upload
	files := form.File["pic"]
	pic := ""

	for _, file := range files {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		uploadedURL, err := utils.UploadFile(f, fileHeader)
		pic = uploadedURL
	}
	// Create a payload for user registration
	fmt.Println(pic)
	payload := &models.RegisterUserSchema{
		Username:     username,
		Email:        email,
		Password:     password,
		PhoneNumber:  phoneNo,
		ProfileImage: pic,
		Role:         role,
	}

	err = services.SignupUser(payload)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Account created successfully. Please verify your account."})
}
