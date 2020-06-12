package request

type OutletCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	SellerID int    `json:"seller_id" validate:"required"`
}

type OutletUpdateRequest struct {
	Name     string `json:"name" validate:"required"`
	SellerID int    `json:"seller_id"`
}
