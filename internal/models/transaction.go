package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	Id     uuid.UUID `db:"id" json:"transaction_id"`
	From   string    `db:"from_user" json:"from_user"`
	Type   string    `db:"type" json:"type"`
	Amount uint64    `db:"amount" json:"amount"`
	To     string    `db:"to_user" json:"to_user"`
	Item   string    `db:"item" json:"item"`
	Date   time.Time `db:"date" json:"date"`
}
