package models

import (
	"time"
)

type StateType uint8

const (
	Active StateType = iota
	Inactive
)

type OutletProduct struct {
	ID         int64
	OutletID   string
	ProductID  string
	Price      int
	OrderPrice int
	State      StateType
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
