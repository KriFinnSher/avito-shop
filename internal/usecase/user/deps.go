package user

import (
	"avito-shop/internal/models"
	"context"
	"github.com/google/uuid"
)

// Repository is an infrastructure layer interface for users
type Repository interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error)
	GetUserByName(ctx context.Context, name string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	UpdateUserBalance(ctx context.Context, userID uuid.UUID, amount int) error
}
