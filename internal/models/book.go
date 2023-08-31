package models

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ISBN        string    `gorm:"varchar(255);uniqueIndex;not null" json:"isbn"`
	Title       string    `gorm:"varchar(255)" json:"title,omitempty"`
	Author      string    `gorm:"varchar(255)" json:"author,omitempty"`
	Description string    `gorm:"text" json:"description,omitempty"`
	Price       float64   `gorm:"type:decimal(10, 2)" json:"price,omitempty"`
	Quantity    int       `gorm:"int" json:"quantity,omitempty"`
	Genre       string    `gorm:"varchar(100)" json:"genre"`
	FullText    string    `gorm:"varchar(255)" json:"full_text"`
	CoverImages string    `gorm:"varchar(255)" json:"cover_images"`
	Sample      string    `gorm:"varchar(255)" json:"sample,omitempty"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
