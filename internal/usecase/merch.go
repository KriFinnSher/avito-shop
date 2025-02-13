package usecase

import (
	"avito-shop/internal/models"
	"avito-shop/internal/usecase/merch"
	"context"
)

type MerchUsecase struct {
	MerchRepo merch.Repository
}

func NewMerchUsecase(repo merch.Repository) *MerchUsecase {
	return &MerchUsecase{MerchRepo: repo}
}

func (u *MerchUsecase) Exist(ctx context.Context, name string) (models.Item, bool) {
	item, err := u.MerchRepo.GetMerch(ctx, name)
	if err != nil {
		return models.Item{}, false
	}
	return item, true
}
