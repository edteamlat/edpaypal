package invoice

import (
	"time"

	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(order *models.Order, subscriptionID uuid.UUID) error {
	i := models.Invoice{}

	if order.IsSubscription {
		i.SubscriptionID = subscriptionID
	}

	i.ID = uuid.New()
	i.Email = order.Email
	i.InvoiceDate = time.Now()
	i.IsProduct = order.IsProduct
	i.IsSubscription = order.IsSubscription
	i.Price = order.Price
	i.ProductID = order.ProductID

	return uc.storage.Create(&i)
}

func (uc UseCase) ByEmail(email string) (models.Invoices, error) {
	return uc.storage.ByEmail(email)
}
