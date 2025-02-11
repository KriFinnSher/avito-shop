package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id       uuid.UUID `json:"transaction_id"`
	UserId   uuid.UUID `json:"user_id"`
	Type     string    `json:"type"`
	Amount   uint64    `json:"amount"`
	TargetId uuid.UUID `json:"target_user_id"`
	ItemId   uuid.UUID `json:"item_id"`
	Date     time.Time `json:"date"`
}
