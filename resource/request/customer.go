package request

type CustomerCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type CustomerUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
