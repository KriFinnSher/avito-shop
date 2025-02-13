package transaction

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

func (r *Repository) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	query := `INSERT INTO transactions (id, from_user, type, amount, to_user, item, date) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, transaction.ID, transaction.From, transaction.Type,
		transaction.Amount, transaction.To, transaction.Item, transaction.Date)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserTransactions(_ context.Context, name string) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, 0)
	query := `SELECT * FROM transactions WHERE from_user = $1 OR to_user = $1`
	err := r.db.Select(&transactions, query, name)
	if err != nil {
		return []models.Transaction{}, err
	}
	return transactions, nil
}

func (r *Repository) GetUserPurchases(_ context.Context, name string) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, 0)
	query := `SELECT * FROM transactions WHERE from_user = $1 AND type = 'purchase'`
	err := r.db.Select(&transactions, query, name)
	if err != nil {
		return []models.Transaction{}, err
	}
	return transactions, nil
}
