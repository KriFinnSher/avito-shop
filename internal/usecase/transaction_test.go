package usecase

import (
	"avito-shop/internal/models"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetUserTransactions(ctx context.Context, name string) ([]models.Transaction, error) {
	args := m.Called(ctx, name)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetUserPurchases(ctx context.Context, name string) ([]models.Transaction, error) {
	args := m.Called(ctx, name)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func TestTransactionUsecase_Send(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	usecase := NewTransactionUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful transfer", func(t *testing.T) {
		fromUser := &models.User{Name: "Alice", Balance: 100}
		toUser := &models.User{Name: "Bob", Balance: 50}
		amount := uint64(30)

		mockRepo.On("CreateTransaction", ctx, mock.Anything).Return(nil)

		err := usecase.Send(ctx, fromUser, toUser, amount)

		assert.NoError(t, err)
		assert.Equal(t, uint64(70), fromUser.Balance)
		assert.Equal(t, uint64(80), toUser.Balance)
		mockRepo.AssertExpectations(t)
	})

	t.Run("insufficient amount", func(t *testing.T) {
		fromUser := &models.User{Name: "Alice", Balance: 100}
		toUser := &models.User{Name: "Bob", Balance: 50}
		amount := uint64(0)

		err := usecase.Send(ctx, fromUser, toUser, amount)

		assert.Error(t, err)
		assert.Equal(t, "insufficient amount", err.Error())
	})
}

func TestTransactionUsecase_Purchase(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	usecase := NewTransactionUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful purchase", func(t *testing.T) {
		user := &models.User{Name: "Alice", Balance: 100}
		item := &models.Item{Name: "T-shirt", Cost: 50}

		mockRepo.On("CreateTransaction", ctx, mock.Anything).Return(nil)

		err := usecase.Purchase(ctx, user, item)

		assert.NoError(t, err)
		assert.Equal(t, uint64(50), user.Balance)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not enough money", func(t *testing.T) {
		user := &models.User{Name: "Alice", Balance: 30}
		item := &models.Item{Name: "Laptop", Cost: 50}

		err := usecase.Purchase(ctx, user, item)

		assert.Error(t, err)
		assert.Equal(t, "not enough money", err.Error())
	})

}

func TestTransactionUsecase_GetHistory(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	usecase := NewTransactionUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful history retrieval", func(t *testing.T) {
		userName := "Alice"
		transactions := []models.Transaction{
			{From: "Alice", To: "Bob", Amount: 30, Type: "transfer", Date: time.Now()},
			{From: "Bob", To: "Alice", Amount: 20, Type: "transfer", Date: time.Now()},
		}

		mockRepo.On("GetUserTransactions", ctx, userName).Return(transactions, nil)

		history, err := usecase.GetHistory(ctx, userName)

		assert.NoError(t, err)
		assert.Len(t, history.Sent, 1)
		assert.Len(t, history.Received, 1)
		mockRepo.AssertExpectations(t)
	})

}

func TestTransactionUsecase_GetInventory(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	usecase := NewTransactionUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful inventory retrieval", func(t *testing.T) {
		userName := "Alice"
		purchases := []models.Transaction{
			{From: "Alice", Item: "Laptop", Amount: 50, Type: "purchase"},
			{From: "Alice", Item: "Phone", Amount: 30, Type: "purchase"},
			{From: "Alice", Item: "Laptop", Amount: 50, Type: "purchase"},
		}

		mockRepo.On("GetUserPurchases", ctx, userName).Return(purchases, nil)

		inventory, err := usecase.GetInventory(ctx, userName)

		assert.NoError(t, err)
		assert.Len(t, inventory, 2)
		assert.Equal(t, uint64(2), inventory[0].Quantity)
		mockRepo.AssertExpectations(t)
	})
}
