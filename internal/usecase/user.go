package usecase

import (
	"avito-shop/internal/models"
	"avito-shop/internal/usecase/user"
	"context"
	"github.com/google/uuid"
)

type UserUsecase struct {
	UserRepo user.UserRepository
}

func NewUserUsecase(repo user.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (u *UserUsecase) GetInventory(ctx context.Context, userID uuid.UUID) (map[string]uint64, error) {
	user_, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return map[string]uint64{}, err // TODO: wrap this error
	}
	items := user_.Items
	inventory := make(map[string]uint64)
	for _, item := range items {
		inventory[item.Name]++
	}
	return inventory, nil
}

func (u *UserUsecase) Exist(ctx context.Context, name string) (models.User, bool) {
	user_, err := u.UserRepo.GetUserByName(ctx, name)
	if err != nil {
		return models.User{}, false
	}
	return user_, true
}

func (u *UserUsecase) UpdateBalance(ctx context.Context, userID uuid.UUID, amount uint64) error {
	err := u.UserRepo.UpdateUserBalance(ctx, userID, int(amount))
	if err != nil {
		return err // TODO: wrap this error
	}
	return nil
}

func (u *UserUsecase) GetBalance(ctx context.Context, userID uuid.UUID) (uint64, error) {
	user_, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return 0, err // TODO: wrap this error
	}
	balance := user_.Balance
	return balance, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user_ models.User) error {
	err := u.UserRepo.CreateUser(ctx, user_)
	if err != nil {
		return err // TODO: wrap this error
	}
	return nil
}
