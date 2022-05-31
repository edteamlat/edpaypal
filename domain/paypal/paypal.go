package paypal

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
)

// PayPal is implemented by the Handler
type PayPal interface {
	ProcessRequest(header http.Header, body []byte) error
}

// PortsOut

type Order interface {
	ByID(ID uuid.UUID) (models.Order, error)
}

type Subscription interface {
	Create(s *models.Subscription) error
}

type Invoice interface {
	Create(order *models.Order, subscriptionID uuid.UUID) error
}
