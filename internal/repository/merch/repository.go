package merch

import (
	"avito-shop/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetMerch(ctx context.Context, name string) (models.Item, error) {
	query := `SELECT * FROM Merch WHERE name=$1`
	item := models.Item{}
	err := r.db.QueryRowContext(ctx, query, name).Scan(&item.Id, &item.Name, &item.Cost)
	if err != nil {
		return models.Item{}, err // TODO: wrap this error
	}
	return item, nil
}
