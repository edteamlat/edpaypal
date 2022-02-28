package invoice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/edteamlat/edpaypal/models"
	"github.com/edteamlat/edpaypal/storage/postgres"
	"github.com/google/uuid"
)

const (
	queryInsert  = "INSERT INTO invoices (id, invoice_date, email, is_product, is_subscription, product_id, subscription_id, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	querySelect  = "SELECT id, invoice_date, email, is_product, is_subscription, product_id, subscription_id, price, created_at, updated_at FROM invoices"
	queryByID    = querySelect + " WHERE id = $1"
	queryByEmail = querySelect + " WHERE email = $1"
)

type Invoice struct {
	db *sql.DB
}

func New(db *sql.DB) Invoice {
	return Invoice{db: db}
}

func (i Invoice) Create(invoice *models.Invoice) error {
	emptyContext := context.Background()
	stmt, err := i.db.PrepareContext(emptyContext, queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	productID := sql.NullString{}
	subscriptionID := sql.NullString{}
	if invoice.ProductID != uuid.Nil {
		productID.String = invoice.ProductID.String()
		productID.Valid = true
	}
	if invoice.SubscriptionID != uuid.Nil {
		subscriptionID.String = invoice.SubscriptionID.String()
		subscriptionID.Valid = true
	}

	row, err := stmt.ExecContext(
		emptyContext,
		invoice.ID,
		invoice.InvoiceDate,
		invoice.Email,
		invoice.IsProduct,
		invoice.IsSubscription,
		productID,
		subscriptionID,
		invoice.Price,
	)
	if err != nil {
		return err
	}

	got, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if got != 1 {
		return fmt.Errorf("expected 1 row, got %d", got)
	}

	return nil
}

func (i Invoice) ByID(ID string) (models.Invoice, error) {
	emptyContext := context.Background()
	stmt, err := i.db.PrepareContext(emptyContext, queryByID)
	if err != nil {
		return models.Invoice{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(emptyContext, ID)
	if err != nil {
		return models.Invoice{}, err
	}

	return i.scan(row)
}

func (i Invoice) ByEmail(email string) (models.Invoices, error) {
	emptyContext := context.Background()
	stmt, err := i.db.PrepareContext(emptyContext, queryByEmail)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(emptyContext, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices models.Invoices
	for rows.Next() {
		invoice, err := i.scan(rows)
		if err != nil {
			return nil, err
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (i Invoice) scan(row postgres.RowScanner) (models.Invoice, error) {
	productID := sql.NullString{}
	subscriptionID := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	invoice := models.Invoice{}

	err := row.Scan(
		&invoice.ID,
		&invoice.InvoiceDate,
		&invoice.Email,
		&invoice.IsProduct,
		&invoice.IsSubscription,
		&productID,
		&subscriptionID,
		&invoice.Price,
		&invoice.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return models.Invoice{}, err
	}

	if productID.String != "" {
		invoice.ProductID = uuid.MustParse(productID.String)
	}
	if subscriptionID.String != "" {
		invoice.SubscriptionID = uuid.MustParse(subscriptionID.String)
	}

	invoice.UpdatedAt = updatedAtNull.Time

	return invoice, nil
}
