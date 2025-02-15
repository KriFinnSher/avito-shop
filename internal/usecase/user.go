package usecase

import (
	"avito-shop/internal/models"
	"avito-shop/internal/usecase/user"
	"context"
	"github.com/google/uuid"
)

// UserUsecase is a struct type for an interaction
// with business-logic layer (user model)
type UserUsecase struct {
	UserRepo user.Repository
}

// NewUserUsecase returns new [UserUsecase]
func NewUserUsecase(repo user.Repository) *UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

// Exist checks and returns user struct if user with defined name is present in db
func (u *UserUsecase) Exist(ctx context.Context, name string) (models.User, bool) {
	user_, err := u.UserRepo.GetUserByName(ctx, name)
	if err != nil {
		return models.User{}, false
	}
	return user_, true
}

// UpdateBalance update user's balance with defined userID
func (u *UserUsecase) UpdateBalance(ctx context.Context, userID uuid.UUID, amount uint64) error {
	err := u.UserRepo.UpdateUserBalance(ctx, userID, int(amount))
	if err != nil {
		return err
	}
	return nil
}

// GetBalance returns user's balance if user with defined userID is present in db
func (u *UserUsecase) GetBalance(ctx context.Context, userID uuid.UUID) (uint64, error) {
	user_, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err
	}
	balance := user_.Balance
	return balance, nil
}

// CreateUser creates new user and ties him with a struct
func (u *UserUsecase) CreateUser(ctx context.Context, user_ models.User) error {
	err := u.UserRepo.CreateUser(ctx, user_)
	if err != nil {
		return err
	}
	return nil
}
