package transaction

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) error
	GetTransactionsByUserId(ctx context.Context, userId uuid.UUID) ([]models.Transaction, error)
}
