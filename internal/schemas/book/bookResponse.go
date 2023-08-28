package book

import (
	"encoding/json"
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
)

type BookResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	CoverImages []string  `json:"cover_images,omitempty"`
}

type BookDetailResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	ISBN        string    `json:"isbn,omitempty"`
	Title       string    `json:"title,omitempty"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	UserID      uuid.UUID `json:"user_id,omitempty"`
	CoverImages []string  `json:"cover_images,omitempty"`
}

func MapBooksToResponse(books []models.Book) []BookResponse {
	bookResponses := make([]BookResponse, len(books))
	for i, book := range books {
		var coverImages []string
		_ = json.Unmarshal([]byte(book.CoverImages), &coverImages)

		bookResponses[i] = BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			Price:       book.Price,
			CoverImages: coverImages,
		}
	}
	return bookResponses
}

func MapBookDetailToResponse(book models.Book) BookDetailResponse {
	var coverImages []string
	_ = json.Unmarshal([]byte(book.CoverImages), &coverImages)

	return BookDetailResponse{
		ID:          book.ID,
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Price:       book.Price,
		CoverImages: coverImages,
		Genre:       book.Genre,
		UserID:      book.UserID,
	}
}
