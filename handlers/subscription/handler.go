package subscription

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/subscription"
	"github.com/edteamlat/edpaypal/models"
)

type Handler struct {
	useCase subscription.Subscription
}

func New(uc subscription.Subscription) Handler {
	return Handler{useCase: uc}
}

func (h Handler) ByEmail(c echo.Context) error {
	email := c.Param("email")

	subscriptions, err := h.useCase.ByEmail(email)
	if err != nil {
		msg := map[string]string{
			"error":    "no se pudo consultar las suscripciones",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]models.Subscriptions{"data": subscriptions}
	return c.JSON(http.StatusOK, msg)
}