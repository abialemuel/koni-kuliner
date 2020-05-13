package request

type OutletProductCreateRequest struct {
	ProductID  int `json:"product_id" validate:"required"`
	OutletID   int `json:"outlet_id" validate:"required"`
	Price      int `json:"price" validate:"required"`
	OrderPrice int `json:"order_price" validate:"required"`
}

type OutletProductUpdateRequest struct {
	State      string `json:"state"`
	Price      int    `json:"price"`
	OrderPrice int    `json:"order_price"`
}
