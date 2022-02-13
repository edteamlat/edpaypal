package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/invoice"
	handler "github.com/edteamlat/edpaypal/handlers/invoice"
	storage "github.com/edteamlat/edpaypal/storage/postgres/invoice"
)

func Invoice(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := invoice.New(store)
	handler.NewRouter(e, useCase)
}
