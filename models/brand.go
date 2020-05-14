package models

import (
	"time"
)

type Brand struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
