package response

import (
	"time"
)

type OutletProductResponse struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	OutletID   int       `json:"outlet_id"`
	Price      int       `json:"price"`
	OrderPrice int       `json:"order_price"`
	State      string    `json:"state"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
