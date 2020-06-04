package request

type SellerCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type SellerUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
