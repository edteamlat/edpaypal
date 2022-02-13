package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/order"
	handler "github.com/edteamlat/edpaypal/handlers/order"
	storage "github.com/edteamlat/edpaypal/storage/postgres/order"
)

func Order(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := order.New(store)
	handler.NewRouter(e, useCase)
}
