package usecase

import (
	"avito-shop/internal/models"
	"avito-shop/internal/usecase/transaction"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity uint64 `json:"quantity"`
}

type receivedCoin struct {
	FromUser string `json:"fromUser"`
	Amount   uint64 `json:"amount"`
}

type sentCoin struct {
	ToUser string `json:"toUser"`
	Amount uint64 `json:"amount"`
}

type CoinHistory struct {
	Received []receivedCoin `json:"received"`
	Sent     []sentCoin     `json:"sent"`
}

type TransactionUsecase struct {
	TransactionRepo transaction.Repository
}

func NewTransactionUsecase(repo transaction.Repository) *TransactionUsecase {
	return &TransactionUsecase{TransactionRepo: repo}
}

func (u *TransactionUsecase) Send(ctx context.Context, fromUser *models.User, toUser *models.User, amount uint64) error {
	if amount <= 0 {
		return errors.New("insufficient amount")
	}
	fromUser.Balance -= amount
	toUser.Balance += amount

	transaction_ := models.Transaction{
		ID:     uuid.New(),
		From:   fromUser.Name,
		Type:   "transfer",
		Amount: amount,
		To:     toUser.Name,
		Item:   "None",
		Date:   time.Now(),
	}

	err := u.TransactionRepo.CreateTransaction(ctx, transaction_)
	if err != nil {
		return err
	}
	fromUser.History = append(fromUser.History, transaction_)
	toUser.History = append(toUser.History, transaction_)

	return nil
}

func (u *TransactionUsecase) Purchase(ctx context.Context, user_ *models.User, item *models.Item) error {
	if item.Cost > user_.Balance {
		return errors.New("not enough money")
	}
	user_.Balance -= item.Cost

	transaction_ := models.Transaction{
		ID:     uuid.New(),
		From:   user_.Name,
		Type:   "purchase",
		Amount: item.Cost,
		To:     "None",
		Item:   item.Name,
		Date:   time.Now(),
	}

	err := u.TransactionRepo.CreateTransaction(ctx, transaction_)
	if err != nil {
		return err
	}
	user_.History = append(user_.History, transaction_)
	user_.Items = append(user_.Items, *item)

	return nil
}

func (u *TransactionUsecase) GetHistory(ctx context.Context, name string) (CoinHistory, error) {
	transactions, err := u.TransactionRepo.GetUserTransactions(ctx, name)
	if err != nil {
		return CoinHistory{}, err
	}

	sentHistory := make([]sentCoin, 0, len(transactions))
	receiveHistory := make([]receivedCoin, 0, len(transactions))
	for _, t := range transactions {
		if t.From == name && t.Type == "transfer" {
			sentHistory = append(sentHistory, sentCoin{
				ToUser: t.To,
				Amount: t.Amount,
			})
		}
		if t.To == name && t.Type == "transfer" {
			receiveHistory = append(receiveHistory, receivedCoin{
				FromUser: t.From,
				Amount:   t.Amount,
			})
		}
	}
	history := CoinHistory{
		Received: receiveHistory,
		Sent:     sentHistory,
	}
	return history, nil
}

func (u *TransactionUsecase) GetInventory(ctx context.Context, name string) ([]InventoryItem, error) {
	purchases, err := u.TransactionRepo.GetUserPurchases(ctx, name)
	if err != nil {
		return nil, err
	}

	itemCounts := make(map[string]uint64)
	for _, purchase := range purchases {
		itemCounts[purchase.Item]++
	}

	inventory := make([]InventoryItem, 0, len(itemCounts))
	for itemType, quantity := range itemCounts {
		inventory = append(inventory, InventoryItem{
			Type:     itemType,
			Quantity: quantity,
		})
	}

	return inventory, nil
}
