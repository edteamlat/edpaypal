package product

import (
	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) All() (models.Products, error) {
	return uc.storage.All()
}

func (uc UseCase) ByID(ID uuid.UUID) (models.Product, error) {
	return uc.storage.ByID(ID)
}
