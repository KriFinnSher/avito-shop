package usecase

import (
	"avito-shop/internal/models"
	"avito-shop/internal/usecase/transaction"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TransactionUsecase struct {
	TransactionRepo transaction.TransactionRepository
}

func NewTransactionUsecase(repo transaction.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{TransactionRepo: repo}
}

func (u *TransactionUsecase) Send(ctx context.Context, fromUser *models.User, toUser *models.User, amount uint64) error {
	if amount <= 0 {
		return errors.New("insufficient amount") // TODO: wrap this error
	}
	fromUser.Balance -= amount
	toUser.Balance += amount

	transaction_ := models.Transaction{
		Id:     uuid.New(),
		From:   fromUser.Name,
		Type:   "transfer",
		Amount: amount,
		To:     toUser.Name,
		Item:   "None",
		Date:   time.Now(),
	}

	err := u.TransactionRepo.CreateTransaction(ctx, transaction_)
	if err != nil {
		return err // TODO: wrap this error
	}
	fromUser.History = append(fromUser.History, transaction_)
	fmt.Println(fromUser.Name, "has history:", fromUser.History)
	toUser.History = append(toUser.History, transaction_)

	return nil
}

func (u *TransactionUsecase) Purchase(ctx context.Context, user_ models.User, item models.Item) error {
	if item.Cost > user_.Balance {
		return errors.New("not enough money") // TODO: wrap this error
	}
	user_.Balance -= item.Cost
	fmt.Println("ITEM:", item)

	transaction_ := models.Transaction{
		Id:     uuid.New(),
		From:   user_.Name,
		Type:   "purchase",
		Amount: item.Cost,
		To:     "None",
		Item:   item.Name,
		Date:   time.Now(),
	}

	err := u.TransactionRepo.CreateTransaction(ctx, transaction_)
	if err != nil {
		return err // TODO: wrap this error
	}
	user_.History = append(user_.History, transaction_)
	user_.Items = append(user_.Items, item)

	return nil
}

func (u *TransactionUsecase) GetHistory(ctx context.Context, name string) ([]models.Transaction, error) {
	transactions, err := u.TransactionRepo.GetUserTransactions(ctx, name)
	if err != nil {
		return nil, err // TODO: wrap this error
	}
	return transactions, nil
}
