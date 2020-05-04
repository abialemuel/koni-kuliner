package request

type OutletCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
