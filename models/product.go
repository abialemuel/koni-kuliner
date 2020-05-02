package models

import (
	"time"
)

type Product struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
