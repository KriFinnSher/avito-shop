package models

import (
	"github.com/google/uuid"
)

type Item struct {
	Id   uuid.UUID `json:"item_id"`
	Name string    `json:"item_name"`
	Cost uint64    `json:"item_cost"`
}
