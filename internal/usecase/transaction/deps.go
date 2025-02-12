package transaction

import (
	"avito-shop/internal/models"
	"context"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	GetUserTransactions(ctx context.Context, name string) ([]models.Transaction, error)
}
