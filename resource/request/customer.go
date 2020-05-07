package request

type CustomerCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"optional"`
	Phone   string `json:"phone" validate:"optional"`
}

type CustomerUpdateRequest struct {
	Name    string `json:"name" validate:"optional"`
	Address string `json:"address" validate:"optional"`
	Phone   string `json:"phone" validate:"optional"`
}
