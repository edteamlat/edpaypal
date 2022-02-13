package subscription

import "github.com/edteamlat/edpaypal/models"

type Storage interface {
	Create(s *models.Subscription) error
	ByEmail(email string) (models.Subscriptions, error)
}

type Subscription interface {
	ByEmail(email string) (models.Subscriptions, error)
}
