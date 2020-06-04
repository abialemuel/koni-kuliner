package models

import (
	"time"
)

type Seller struct {
	ID        int64
	Name      string
	Username  string
	Password  string
	Address   string
	Phone     string
	Outlets   []Outlet
	CreatedAt time.Time
	UpdatedAt time.Time
}
