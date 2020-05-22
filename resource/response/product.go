package response

import (
	"time"
)

type ProductResponse struct {
	ID        int64               `json:"id"`
	Name      string              `json:"name"`
	Brand     DetailBrandResponse `json:"brand"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

type DetailProductResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
