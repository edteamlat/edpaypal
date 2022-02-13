package product

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/product"
	"github.com/edteamlat/edpaypal/models"
)

type Handler struct {
	useCase product.Product
}

func New(uc product.Product) Handler {
	return Handler{useCase: uc}
}

func (h Handler) All(c echo.Context) error {
	data, err := h.useCase.All()
	if err != nil {
		msg := map[string]string{
			"error":    "no pudimos consultar la info",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]models.Products{"data": data}
	return c.JSON(http.StatusOK, msg)
}

func (h Handler) ByID(c echo.Context) error {
	ID := c.Param("id")
	data, err := h.useCase.ByID(uuid.MustParse(ID))
	if err != nil {
		msg := map[string]string{
			"error":    "No se pudo consultar el producto",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]models.Product{"data": data}
	return c.JSON(http.StatusOK, msg)
}