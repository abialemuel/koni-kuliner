package response

import (
	"time"
)

type TransactionResponse struct {
	ID        int64                    `json:"id"`
	Customer  DetailCustomerResponse   `json:"customer"`
	CartItems []DetailCartItemResponse `json:"cart_items"`
	Amount    int                      `json:"amount"`
	State     string                   `json:"state"`
	Delivery  string                   `json:"delivery"`
	Note      string                   `json:"note"`
	PoDate    time.Time                `json:"po_date"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type TransactionDetailResponse struct {
	ID        int64                    `json:"id"`
	Customer  DetailCustomerResponse   `json:"customer"`
	CartItems []DetailCartItemResponse `json:"cart_items"`
	Amount    int                      `json:"amount"`
	State     string                   `json:"state"`
	Delivery  string                   `json:"delivery"`
	Note      string                   `json:"note"`
	Feedback  string                   `json:"feedback"`
	PoDate    time.Time                `json:"po_date"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type DetailCustomerResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type DetailCartItemResponse struct {
	ID       int64                 `json:"id"`
	Product  DetailProductResponse `json:"product"`
	Quantity int                   `json:"quantity"`
	Price    int                   `json:"price"`
}

type DetailProductResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
