package usecase

import (
	"avito-shop/internal/models"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockMerchRepository struct {
	mock.Mock
}

func (m *MockMerchRepository) GetMerch(ctx context.Context, name string) (models.Item, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models.Item), args.Error(1)
}

func TestMerchUsecase_Exist(t *testing.T) {
	mockRepo := new(MockMerchRepository)
	usecase := NewMerchUsecase(mockRepo)
	ctx := context.Background()

	t.Run("merch exists", func(t *testing.T) {
		item := models.Item{Name: "T-Shirt"}
		mockRepo.On("GetMerch", ctx, "T-Shirt").Return(item, nil)

		result, exists := usecase.Exist(ctx, "T-Shirt")

		assert.True(t, exists)
		assert.Equal(t, item, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("merch does not exist", func(t *testing.T) {
		mockRepo.On("GetMerch", ctx, "NonExistent").Return(models.Item{}, errors.New("not found"))

		result, exists := usecase.Exist(ctx, "NonExistent")

		assert.False(t, exists)
		assert.Equal(t, models.Item{}, result)
		mockRepo.AssertExpectations(t)
	})
	t.Run("empty name", func(t *testing.T) {
		mockRepo.On("GetMerch", ctx, "").Return(models.Item{}, errors.New("invalid name"))

		result, exists := usecase.Exist(ctx, "")

		assert.False(t, exists)
		assert.Equal(t, models.Item{}, result)
		mockRepo.AssertExpectations(t)
	})
}
