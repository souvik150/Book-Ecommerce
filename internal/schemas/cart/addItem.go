package cart

import "github.com/google/uuid"

type AddItemToCartSchema struct {
	BookID   uuid.UUID `json:"book_id"`
	Quantity int       `json:"quantity"`
}
