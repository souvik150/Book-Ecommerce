package services

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

func CreateBook(payload *models.CreateBookSchema) (*models.Book, error) {

	// Convert string price to float64
	price, err := strconv.ParseFloat(payload.Price, 64)
	if err != nil {
		return nil, err
	}

	// Convert string quantity to integer
	quantity, err := strconv.Atoi(payload.Quantity)
	if err != nil {
		return nil, err
	}

	coverImagesJSON, err := json.Marshal(payload.CoverImages)
	if err != nil {
		return nil, err
	}

	// Create a new book
	book := &models.Book{
		ISBN:        payload.ISBN,
		Title:       payload.Title,
		Author:      payload.Author,
		Description: payload.Description,
		Genre:       payload.Genre,
		Price:       price,
		Quantity:    quantity,
		FullText:    payload.FullText,
		Sample:      payload.Sample,
		CoverImages: string(coverImagesJSON),
		UserID:      payload.UserID,
	}

	// Store the book in the database
	result := database.DB.Create(&book)
	if result.Error != nil {
		return nil, err
	}

	return book, nil
}

func GetBooksPaginated(page, limit int) ([]models.Book, error) {
	offset := (page - 1) * limit

	var books []models.Book
	result := database.DB.Offset(offset).Limit(limit).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func GetBookByID(bookID uuid.UUID) (*models.Book, error) {
	var book models.Book
	if err := database.DB.First(&book, "id = ?", bookID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return &book, nil
}

func GetBookByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func UpdateBook(book *models.Book) (*models.Book, error) {
	result := database.DB.Save(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func DeleteBook(userID uuid.UUID, bookID string) error {
	var book models.Book

	result := database.DB.First(&book, "id = ?", bookID)
	if result.Error != nil {
		return result.Error
	}

	if userID != book.UserID {
		return gorm.ErrRecordNotFound // Return a specific error indicating unauthorized
	}

	result = database.DB.Delete(&book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
