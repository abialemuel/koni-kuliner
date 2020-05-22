package response

import (
	"time"
)

type CartItemResponse struct {
	ID              int64                 `json:"id"`
	OutletProductID int                   `json:"outlet_product_id"`
	Product         DetailProductResponse `json:"product"`
	Price           int                   `json:"price"`
	OrderPrice      int                   `json:"order_price"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
}

type DetailCartItemResponse struct {
	ID              int64                 `json:"id"`
	OutletProductID int                   `json:"outlet_product_id"`
	Product         DetailProductResponse `json:"product"`
	Quantity        int                   `json:"quantity"`
	Price           int                   `json:"price"`
	Total           int                   `json:"total"`
}
