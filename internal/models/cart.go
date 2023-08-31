package models

import (
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	UserID       uuid.UUID  `gorm:"type:uuid" json:"user_id,omitempty"`
	CreatedAt    time.Time  `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt    time.Time  `gorm:"not null" json:"updated_at,omitempty"`
	Wishlist     bool       `gorm:"not null;default:false" json:"wishlist,omitempty"`
	SaveForLater bool       `gorm:"not null;default:false" json:"save_for_later,omitempty"`
	Items        []CartItem `json:"items,omitempty"`
	TotalCost    float64    `gorm:"type:decimal(10, 2)" json:"total_cost,omitempty"`
	Active       bool       `gorm:"not null;default:true" json:"active"`
}

type CartItem struct {
	CartItemID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"cart_item_id,omitempty"`
	BookID     uuid.UUID `gorm:"type:uuid" json:"book_id,omitempty"`
	CartID     uuid.UUID `gorm:"type:uuid" json:"cart_id,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
}
