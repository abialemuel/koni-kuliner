package models

import (
	"time"
)

type CartItem struct {
	ID              int64
	CustomerID      int
	OutletProductID int
	TransactionID   int
	Price           int
	OrderPrice      int
	Quantity        int
	Customer        Customer
	OutletProduct   OutletProduct
	Transaction     Transaction
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
