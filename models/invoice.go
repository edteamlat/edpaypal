package models

import (
	"time"

	"github.com/google/uuid"
)

// Invoice structure
type Invoice struct {
	ID             uuid.UUID `json:"id"`
	InvoiceDate    time.Time `json:"invoice_date"`
	Email          string    `json:"email"`
	IsProduct      bool      `json:"is_product"`
	IsSubscription bool      `json:"is_subscription"`
	ProductID      uuid.UUID `json:"product_id"`
	SubscriptionID uuid.UUID `json:"subscription_id"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Invoices slice of invoice
type Invoices []Invoice
