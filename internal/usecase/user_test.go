package usecase

import (
	"avito-shop/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByName(ctx context.Context, name string) (models.User, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserBalance(ctx context.Context, userID uuid.UUID, amount int) error {
	args := m.Called(ctx, userID, amount)
	return args.Error(0)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user_ models.User) error {
	args := m.Called(ctx, user_)
	return args.Error(0)
}

func TestUserUsecase_Exist(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)
	ctx := context.Background()

	t.Run("user exists", func(t *testing.T) {
		userID := uuid.New()
		user := models.User{Name: "Alice", ID: userID}

		mockRepo.On("GetUserByName", ctx, "Alice").Return(user, nil)

		result, found := usecase.Exist(ctx, "Alice")
		assert.True(t, found)
		assert.Equal(t, user, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("user does not exist", func(t *testing.T) {
		mockRepo.On("GetUserByName", ctx, "NonExistentUser").Return(models.User{}, errors.New("user not found"))

		result, found := usecase.Exist(ctx, "NonExistentUser")
		assert.False(t, found)
		assert.Equal(t, models.User{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_UpdateBalance(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful balance update", func(t *testing.T) {
		userID := uuid.New()
		amount := uint64(100)

		mockRepo.On("UpdateUserBalance", ctx, userID, int(amount)).Return(nil)

		err := usecase.UpdateBalance(ctx, userID, amount)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		userID := uuid.New()
		amount := uint64(100)

		mockRepo.On("UpdateUserBalance", ctx, userID, int(amount)).Return(errors.New("repo error"))

		err := usecase.UpdateBalance(ctx, userID, amount)
		assert.Error(t, err)
		assert.Equal(t, "repo error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_GetBalance(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful balance retrieval", func(t *testing.T) {
		userID := uuid.New()
		user := models.User{ID: userID, Balance: 100}

		mockRepo.On("GetUserByID", ctx, userID).Return(user, nil)

		balance, err := usecase.GetBalance(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, uint64(100), balance)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		userID := uuid.New()

		mockRepo.On("GetUserByID", ctx, userID).Return(models.User{}, errors.New("repo error"))

		balance, err := usecase.GetBalance(ctx, userID)
		assert.Error(t, err)
		assert.Equal(t, "repo error", err.Error())
		assert.Equal(t, uint64(0), balance)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)
	ctx := context.Background()

	t.Run("successful user creation", func(t *testing.T) {
		userID := uuid.New()
		user := models.User{Name: "Alice", ID: userID, Balance: 100}

		mockRepo.On("CreateUser", ctx, user).Return(nil)

		err := usecase.CreateUser(ctx, user)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("repository error", func(t *testing.T) {
		userID := uuid.New()
		user := models.User{Name: "Alice", ID: userID, Balance: 100}

		mockRepo.On("CreateUser", ctx, user).Return(errors.New("repo error"))

		err := usecase.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Equal(t, "repo error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
