package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Image          string    `json:"image"`
	IsSubscription bool      `json:"is_subscription"`
	Months         uint8     `json:"months"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Products slice of product
type Products []Product
