package transaction

import (
	"avito-shop/internal/models"
	"context"
)

// Repository is an infrastructure layer interface for transactions
type Repository interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	GetUserTransactions(ctx context.Context, name string) ([]models.Transaction, error)
	GetUserPurchases(ctx context.Context, name string) ([]models.Transaction, error)
}
