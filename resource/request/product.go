package request

type ProductCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	BrandID  int    `json:"brand_id" validate:"required"`
	SellerID int    `json:"seller_id" validate:"required"`
}

type ProductUpdateRequest struct {
	Name     string `json:"name"`
	BrandID  int    `json:"brand_id"`
	SellerID int    `json:"seller_id"`
}
