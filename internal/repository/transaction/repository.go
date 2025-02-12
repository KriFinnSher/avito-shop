package transaction

import (
	"avito-shop/internal/models"
	"context"
	"fmt"
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
	_, err := r.db.ExecContext(ctx, query, transaction.Id, transaction.From, transaction.Type,
		transaction.Amount, transaction.To, transaction.Item, transaction.Date)
	fmt.Println(err)
	if err != nil {
		return err // TODO: wrap this error
	}
	return nil
}

func (r *Repository) GetUserTransactions(ctx context.Context, name string) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, 0)
	query := `SELECT * FROM transactions WHERE from_user = $1 OR to_user = $1`
	err := r.db.Select(&transactions, query, name)
	fmt.Println(transactions, name, "BIG BOY [DB LEVEL]")
	if err != nil {
		return []models.Transaction{}, err // TODO: wrap this error
	}
	return transactions, nil
}
