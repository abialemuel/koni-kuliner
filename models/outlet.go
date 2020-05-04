package models

import (
	"time"
)

type Outlet struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
