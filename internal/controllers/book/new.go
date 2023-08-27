package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/services"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/utils"
)

func CreateBook(c *fiber.Ctx) error {
	if len(c.Request().Body()) > 1*1024*1024 {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"status":  "fail",
			"message": "Request body size is too large",
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Failed to process form data"})
	}

	// Get the user ID from the authenticated user (adjust according to your authentication logic)
	userID := c.Locals("userID").(uuid.UUID)

	// Get book data from the form
	isbn := form.Value["isbn"][0]

	// If book with same isbn exists, return error
	_, err = services.GetBookByISBN(isbn)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Book with same ISBN already exists"})
	}

	title := form.Value["title"][0]
	author := form.Value["author"][0]
	description := form.Value["description"][0]
	genre := form.Value["genre"][0]
	price := form.Value["price"][0]
	quantity := form.Value["quantity"][0]

	if isbn == "" || title == "" || author == "" || description == "" || genre == "" || price == "" || quantity == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request body"})
	}

	full_text_file := form.File["full_text"]
	sample_file := form.File["sample"]
	cover_files := form.File["cover_images"]

	full_text_url := ""
	sample_url := ""
	cover_images := []string{}

	for _, file := range full_text_file {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to open file"})
		}
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		uploadedURL, err := utils.UploadFile(f, fileHeader)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to upload file"})
		}
		full_text_url = uploadedURL
	}

	fmt.Println(full_text_url)

	for _, file := range sample_file {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to open file"})
		}
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		uploadedURL, err := utils.UploadFile(f, fileHeader)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to upload file"})
		}
		sample_url = uploadedURL
	}

	fmt.Println(sample_url)

	for _, file := range cover_files {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to open file"})
		}
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)

		uploadedURL, err := utils.UploadFile(f, fileHeader)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to upload file"})
		}
		cover_images = append(cover_images, uploadedURL)
	}

	if full_text_url == "" || sample_url == "" || len(cover_images) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Unable to upload files"})
	}

	payload := &models.CreateBookSchema{
		ISBN:        isbn,
		Title:       title,
		Author:      author,
		Description: description,
		Genre:       genre,
		Price:       price,
		Quantity:    quantity,
		FullText:    full_text_url,
		Sample:      sample_url,
		CoverImages: cover_images,
		UserID:      userID,
	}

	book, err := services.CreateBook(payload)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"book": book}})
}
