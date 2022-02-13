package product

import (
	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/product"
)

const (
	path     = "/v1/products"
	pathAll  = ""
	pathByID = "/:id"
)

func NewRouter(e *echo.Echo, useCase product.Product) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathAll, handler.All)
	g.GET(pathByID, handler.ByID)
}