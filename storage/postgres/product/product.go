package product

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/edteamlat/edpaypal/models"
	"github.com/edteamlat/edpaypal/storage/postgres"
)

const (
	query     = "SELECT id, name, description, image, is_subscription, months, price, created_at, updated_at FROM products"
	queryAll  = query + " ORDER BY name"
	queryByID = query + " WHERE id = $1"
)

type Product struct {
	db *sql.DB
}

func New(db *sql.DB) Product {
	return Product{db: db}
}

func (p Product) All() (models.Products, error) {
	emptyContext := context.Background()
	stmt, err := p.db.PrepareContext(emptyContext, queryAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(emptyContext)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp models.Products
	for rows.Next() {
		prod, err := p.scan(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, prod)
	}

	return resp, nil
}

func (p Product) ByID(ID uuid.UUID) (models.Product, error) {
	emptyContext := context.Background()
	stmt, err := p.db.PrepareContext(emptyContext, queryByID)
	if err != nil {
		return models.Product{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(emptyContext, ID)

	return p.scan(row)
}

func (p Product) scan(r postgres.RowScanner) (models.Product, error) {
	updatedNull := sql.NullTime{}
	resp := models.Product{}

	err := r.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.Image,
		&resp.IsSubscription,
		&resp.Months,
		&resp.Price,
		&resp.CreatedAt,
		&updatedNull,
	)
	if err != nil {
		return models.Product{}, err
	}

	resp.UpdatedAt = updatedNull.Time

	return resp, nil
}
