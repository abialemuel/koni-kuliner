package request

type BrandCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type BrandUpdateRequest struct {
	Name string `json:"name" validate:"required"`
}
