package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/subscription"
	handler "github.com/edteamlat/edpaypal/handlers/subscription"
	storage "github.com/edteamlat/edpaypal/storage/postgres/subscription"
)

func Subscription(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := subscription.New(store)
	handler.NewRouter(e, useCase)
}
