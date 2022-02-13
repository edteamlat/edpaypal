package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/edteamlat/edpaypal/domain/invoice"
	"github.com/edteamlat/edpaypal/domain/order"
	"github.com/edteamlat/edpaypal/domain/paypal"
	"github.com/edteamlat/edpaypal/domain/subscription"

	storageInvoice "github.com/edteamlat/edpaypal/storage/postgres/invoice"
	storageOrder "github.com/edteamlat/edpaypal/storage/postgres/order"
	storageSubscription "github.com/edteamlat/edpaypal/storage/postgres/subscription"

	handler "github.com/edteamlat/edpaypal/handlers/paypal"
)

func PayPal(e *echo.Echo, db *sql.DB) {
	useCaseOrder := buildOrder(db)
	useCaseSubs := buildSubs(db)
	useCaseInvoice := buildInvoice(db)
	useCasePayPal := paypal.New(useCaseOrder, useCaseSubs, useCaseInvoice)

	handler.NewRouter(e, useCasePayPal)
}

func buildOrder(db *sql.DB) paypal.Order {
	store := storageOrder.New(db)
	return order.New(store)
}

func buildSubs(db *sql.DB) paypal.Subscription {
	store := storageSubscription.New(db)
	return subscription.New(store)
}

func buildInvoice(db *sql.DB) paypal.Invoice {
	store := storageInvoice.New(db)
	return invoice.New(store)
}
