package models

import (
	"time"

	"github.com/google/uuid"
)

// Order structure
type Order struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	IsProduct      bool      `json:"is_product"`
	IsSubscription bool      `json:"is_subscription"`
	ProductID      uuid.UUID `json:"product_id"`
	TypeSubs       string    `json:"type_subs"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Orders slice of order
type Orders []Order
