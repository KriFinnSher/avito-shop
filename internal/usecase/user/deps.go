package user

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userId uuid.UUID) (models.User, error)
	GetUserByName(ctx context.Context, name string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	UpdateUserBalance(ctx context.Context, userId uuid.UUID, amount int) error
}
