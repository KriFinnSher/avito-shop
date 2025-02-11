package user

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUser(ctx context.Context, userId uuid.UUID) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	GetBalance(ctx context.Context, userID uuid.UUID) (uint64, error)
	GetHistory(ctx context.Context, userId uuid.UUID) ([]models.Transaction, error)
	GetInventory(ctx context.Context, userId uuid.UUID) (map[models.Item]uint64, error)
}
