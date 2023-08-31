package book

import "github.com/google/uuid"

type CreateBookSchema struct {
	ISBN        string    `json:"isbn"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	Price       string    `json:"price"`
	Quantity    string    `json:"quantity"`
	FullText    string    `json:"full_text_url"`
	Sample      string    `json:"sample_url"`
	CoverImages []string  `json:"cover_images"`
	UserID      uuid.UUID `json:"user_id"`
}
