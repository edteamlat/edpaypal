package subscription

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/edteamlat/edpaypal/models"
	"github.com/edteamlat/edpaypal/storage/postgres"
)

const (
	queryInsert        = "INSERT INTO subscriptions (id, email, status, type_subs, begins_at, ends_at) VALUES ($1, $2, $3, $4, $5, $6)"
	querySelect        = "SELECT id, email, status, type_subs, begins_at, ends_at, created_at, updated_at FROM subscriptions"
	querySelectByEmail = querySelect + " WHERE email = $1"
)

type Subscription struct {
	db *sql.DB
}

func New(db *sql.DB) Subscription {
	return Subscription{db: db}
}

func (s Subscription) Create(subs *models.Subscription) error {
	emptyContext := context.Background()

	stmt, err := s.db.PrepareContext(emptyContext, queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.ExecContext(
		emptyContext,
		subs.ID,
		subs.Email,
		subs.Status,
		subs.TypeSubs,
		subs.BeginsAt,
		subs.EndsAt,
	)
	if err != nil {
		return err
	}

	got, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if got != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", got)
	}

	return nil
}

func (s Subscription) ByEmail(email string) (models.Subscriptions, error) {
	emptyContext := context.Background()

	stmt, err := s.db.PrepareContext(emptyContext, querySelectByEmail)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(
		emptyContext,
		email,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp models.Subscriptions
	for rows.Next() {
		row, err := s.scan(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, row)
	}

	return resp, nil
}

func (s Subscription) scan(r postgres.RowScanner) (models.Subscription, error) {
	updatedAtNull := sql.NullTime{}
	m := models.Subscription{}

	err := r.Scan(
		&m.ID,
		&m.Email,
		&m.Status,
		&m.TypeSubs,
		&m.BeginsAt,
		&m.EndsAt,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return models.Subscription{}, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
