package paypal

import (
	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/paypal"
)

const (
	path       = "/v1/paypal"
	pathCreate = "/webhook"
)

func NewRouter(e *echo.Echo, useCase paypal.PayPal) {
	handler := New(useCase)

	g := e.Group(path)
	g.POST(pathCreate, handler.Webhook)
}
