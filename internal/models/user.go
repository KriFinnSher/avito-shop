package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID     `json:"user_id"`
	Name    string        `json:"username"`
	Hash    string        `json:"password_hash"`
	Balance uint64        `json:"balance"`
	Items   []Item        `json:"inventory"`
	History []Transaction `json:"history"`
}
