package response

import "time"

type OutletResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	SellerID  int       `json:"seller_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
