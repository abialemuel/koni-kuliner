package response

import (
	"time"

	"github.com/koni-kuliner/models"
)

type TransactionResponse struct {
	ID        int64                  `json:"id"`
	Customer  DetailCustomerResponse `json:"customer"`
	CartItems []models.CartItem      `json:"cart_items"`
	Amount    int                    `json:"amount"`
	State     string                 `json:"state"`
	Delivery  string                 `json:"delivery"`
	Note      string                 `json:"note"`
	Feedback  string                 `json:"feedback"`
	PoDate    time.Time              `json:"po_date"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type DetailCustomerResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
