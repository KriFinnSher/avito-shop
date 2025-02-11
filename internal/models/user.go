package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID       `json:"user_id"`
	Name      string          `json:"user_name"`
	Balance   uint64          `json:"balance"`
	Inventory map[Item]uint64 `json:"inventory"`
	History   []Transaction   `json:"history"`
}
