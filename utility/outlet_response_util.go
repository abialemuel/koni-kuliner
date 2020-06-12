package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func OutletResponse(outlet []models.Outlet) []response.OutletResponse {
	var returnedResponse []response.OutletResponse
	for _, outlet := range outlet {
		singleResponse := response.OutletResponse{
			ID:        outlet.ID,
			Name:      outlet.Name,
			SellerID:  outlet.SellerID,
			CreatedAt: outlet.CreatedAt,
			UpdatedAt: outlet.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func OutletDetailResponse(outlet models.Outlet) response.OutletResponse {
	var returnedResponse response.OutletResponse
	singleResponse := response.OutletResponse{
		ID:        outlet.ID,
		Name:      outlet.Name,
		SellerID:  outlet.SellerID,
		CreatedAt: outlet.CreatedAt,
		UpdatedAt: outlet.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
