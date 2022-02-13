package order

import (
	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
)

type Storage interface {
	Create(o *models.Order) error
	ByID(ID uuid.UUID) (models.Order, error)
}

type Order interface {
	Create(o *models.Order) error
	ByID(ID uuid.UUID) (models.Order, error)
}
