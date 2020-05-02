package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
