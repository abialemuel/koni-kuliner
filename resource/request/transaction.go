package request

import "time"

type TransactionCreateRequest struct {
	CartItemIDS []int     `json:"cart_item_ids" validate:"required"`
	PoDate      time.Time `json:"po_date" validate:"required"`
	Delivery    string    `json:"delivery"`
	Note        string    `json:"note"`
}

type TransactionUpdateRequest struct {
	State    int    `json:"state"`
	Delivery int    `json:"delivery"`
	Note     string `json:"note"`
}
