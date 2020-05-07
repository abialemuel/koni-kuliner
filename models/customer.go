package models

import (
	"time"
)

type Customer struct {
	ID        int64
	Name      string
	Address   string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
