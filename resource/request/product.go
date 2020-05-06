package request

type ProductCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type ProductUpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
