package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/product"
	handler "github.com/edteamlat/edpaypal/handlers/product"
	storage "github.com/edteamlat/edpaypal/storage/postgres/product"
)

func Product(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := product.New(store)
	handler.NewRouter(e, useCase)
}
