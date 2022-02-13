package product

import (
	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
)

// ******************
// **** PORT OUT ****
// ******************

type Storage interface {
	All() (models.Products, error)
	ByID(ID uuid.UUID) (models.Product, error)
}

// *****************
// **** PORT IN ****
// *****************

type Product interface {
	All() (models.Products, error)
	ByID(ID uuid.UUID) (models.Product, error)
}
