package cart

import "github.com/google/uuid"

type CreateCartSchema struct {
	ID     uuid.UUID `json:"id,omitempty"`
	UserID uuid.UUID `json:"user_id,omitempty"`
	Items  []uuid.UUID
}
