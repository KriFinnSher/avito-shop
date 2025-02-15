package models

import (
	"github.com/google/uuid"
)

// Item is a representation of Avito merch
type Item struct {
	Id   uuid.UUID `json:"item_id"`
	Name string    `json:"item_name"`
	Cost uint64    `json:"item_cost"`
}
