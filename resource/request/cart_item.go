package request

type CartItemCreateRequest struct {
	OutletProductID int `json:"outlet_product_id" validate:"required"`
	CustomerID      int `json:"customer_id" validate:"required"`
	Quantity        int `json:"quantity" validate:"required"`
}

type CartItemUpdateRequest struct {
	Quantity int `json:"quantity" validate:"required"`
}
