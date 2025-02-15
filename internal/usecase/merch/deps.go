package merch

import (
	"avito-shop/internal/models"
	"context"
)

// Repository is an infrastructure layer interface for merch
type Repository interface {
	GetMerch(ctx context.Context, name string) (models.Item, error)
}
