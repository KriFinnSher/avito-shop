package models

import (
	"github.com/google/uuid"
)

// User is a representation of Avito worker
type User struct {
	ID      uuid.UUID `json:"user_id"`
	Name    string    `json:"username"`
	Hash    string    `json:"password_hash"`
	Balance uint64    `json:"balance"`
}
