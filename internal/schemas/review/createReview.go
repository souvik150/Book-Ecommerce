package review

import "github.com/google/uuid"

type CreateReviewSchema struct {
	UserID  uuid.UUID `json:"user_id"`
	BookID  uuid.UUID `json:"book_id"`
	Comment string    `json:"comment"`
	Rating  float64   `json:"rating"`
}
