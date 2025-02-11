package usecase

import "avito-shop/internal/usecase/user"

type UserUsecase struct {
	UserRepo user.UserRepository
}
