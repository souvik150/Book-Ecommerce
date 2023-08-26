package models

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ISBN          string    `gorm:"varchar(255);uniqueIndex;not null" json:"isbn"`
	Title         string    `gorm:"varchar(255)" json:"title,omitempty"`
	Description   string    `gorm:"text" json:"description,omitempty"`
	Price         float64   `gorm:"type:decimal(10, 2)" json:"price,omitempty"`
	StockQuantity int       `gorm:"int" json:"stock_quantity,omitempty"`
	Genre         string    `gorm:"varchar(100)" json:"genre,omitempty"`
	FullTextURL   string    `gorm:"varchar(255)" json:"full_text_url,omitempty"`
	CoverImages   []string  `gorm:"type:varchar(255)[]" json:"cover_images,omitempty"`
	SampleURL     string    `gorm:"varchar(255)" json:"sample_url,omitempty"`
	UserID        uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
