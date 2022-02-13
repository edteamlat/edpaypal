package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/invoice"
	"github.com/edteamlat/edpaypal/models"
)

type Handler struct {
	useCase invoice.Invoice
}

func New(uc invoice.Invoice) Handler {
	return Handler{useCase: uc}
}

func (h Handler) ByEmail(c echo.Context) error {
	email := c.Param("email")

	invoices, err := h.useCase.ByEmail(email)
	if err != nil {
		msg := map[string]string{
			"error":    "no se pudo consultar las facturas",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]models.Invoices{"data": invoices}
	return c.JSON(http.StatusOK, msg)
}
