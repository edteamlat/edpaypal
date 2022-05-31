package paypal

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/paypal"
)

type Handler struct {
	useCase paypal.PayPal
}

func New(paypal paypal.PayPal) Handler {
	return Handler{useCase: paypal}
}

func (h Handler) Webhook(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		response := map[string]string{"error": "no se pudo leer la información", "internal": err.Error()}
		log.Print(response)
		return c.JSON(http.StatusBadRequest, response)
	}

	go func() {
		c.Logger().Print("Paypal request")
		c.Logger().Print(string(body))
		err = h.useCase.ProcessRequest(c.Request().Header, body)
		if err != nil {
			log.Print("error procesando el webhook", err)
		}
	}()

	msg := map[string]string{"message": "ok"}
	return c.JSON(http.StatusOK, msg)
}
