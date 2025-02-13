package merch

import (
	"avito-shop/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type PostgreRepository struct {
	db *sqlx.DB
}

func NewPostgreRepo(db *sqlx.DB) *PostgreRepository {
	return &PostgreRepository{db: db}
}

func (r *PostgreRepository) GetMerch(ctx context.Context, name string) (models.Item, error) {
	query := `SELECT * FROM Merch WHERE name=$1`
	item := models.Item{}
	err := r.db.QueryRowContext(ctx, query, name).Scan(&item.Id, &item.Name, &item.Cost)
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}
