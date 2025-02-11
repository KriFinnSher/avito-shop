package usecase

import "avito-shop/internal/usecase/transaction"

type TransactionUsecase struct {
	TransactionRepo transaction.TransactionRepository
}
