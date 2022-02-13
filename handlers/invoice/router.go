package invoice

import (
	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/invoice"
)

const (
	path        = "/v1/invoices"
	pathByEmail = "/:email"
)

func NewRouter(e *echo.Echo, useCase invoice.Invoice) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathByEmail, handler.ByEmail)
}
