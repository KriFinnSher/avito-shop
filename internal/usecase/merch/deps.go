package merch

import (
	"avito-shop/internal/models"
	"context"
)

type MerchRepository interface {
	GetMerch(ctx context.Context, name string) (models.Item, error)
}
