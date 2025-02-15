package usecase

import (
	"avito-shop/internal/models"
	"avito-shop/internal/usecase/merch"
	"context"
)

// MerchUsecase is a struct type for an interaction
// with business-logic layer (merch model)
type MerchUsecase struct {
	MerchRepo merch.Repository
}

// NewMerchUsecase returns new [MerchUsecase]
func NewMerchUsecase(repo merch.Repository) *MerchUsecase {
	return &MerchUsecase{MerchRepo: repo}
}

// Exist checks and returns item struct if item with defined name is present in db
func (u *MerchUsecase) Exist(ctx context.Context, name string) (models.Item, bool) {
	item, err := u.MerchRepo.GetMerch(ctx, name)
	if err != nil {
		return models.Item{}, false
	}
	return item, true
}
