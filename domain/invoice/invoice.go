package invoice

import "github.com/edteamlat/edpaypal/models"

type Storage interface {
	Create(i *models.Invoice) error
	ByEmail(email string) (models.Invoices, error)
}

type Invoice interface {
	ByEmail(email string) (models.Invoices, error)
}
