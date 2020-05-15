package models

import (
	"time"
)

type TransactionStateType uint8
type TransactionDeliveryType uint8

const (
	TransactionStatePending TransactionStateType = iota + 1
	TransactionStateProcessed
	TransactionStatePrepared
	TransactionStateDelivered
)

const (
	TransactionDeliveryPickUp TransactionDeliveryType = iota + 1
	TransactionDeliveryShipping
)

type Transaction struct {
	ID       int64
	UserID   int
	Amount   int
	Note     string
	Feedback string
	// CartItems []CartItem
	State     TransactionStateType
	PoDate    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s TransactionStateType) ToString() string {
	types := map[TransactionStateType]string{
		TransactionStatePending:   "pending",
		TransactionStateProcessed: "processed",
		TransactionStatePrepared:  "prepared",
		TransactionStateDelivered: "delivered",
	}
	return types[s]
}
