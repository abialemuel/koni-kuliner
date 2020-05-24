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
	ID         int64
	CustomerID int
	Amount     int
	Note       string
	Feedback   string
	Customer   Customer
	CartItems  []CartItem
	State      TransactionStateType
	Delivery   TransactionDeliveryType
	PoDate     time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
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

func (s TransactionDeliveryType) ToString() string {
	types := map[TransactionDeliveryType]string{
		TransactionDeliveryPickUp:   "pick_up",
		TransactionDeliveryShipping: "shipping",
	}
	return types[s]
}

func ToTransactionStateType(str string) TransactionStateType {
	var transaction TransactionStateType

	switch str {
	case "pending":
		transaction = TransactionStatePending
	case "processed":
		transaction = TransactionStateProcessed
	case "prepared":
		transaction = TransactionStatePrepared
	case "delivered":
		transaction = TransactionStateDelivered
	}
	return transaction
}

func ToTransactionDeliveryType(str string) TransactionDeliveryType {
	var transaction TransactionDeliveryType

	switch str {
	case "pick_up":
		transaction = TransactionDeliveryPickUp
	case "shipping":
		transaction = TransactionDeliveryShipping
	}
	return transaction
}
