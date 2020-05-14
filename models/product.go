package models

import (
	"time"
)

type Product struct {
	ID        int64
	Name      string
	BrandID   int
	Brand     Brand
	CreatedAt time.Time
	UpdatedAt time.Time
}
