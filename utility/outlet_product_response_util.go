package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func OutletProductResponse(outletPoduct []models.OutletProduct) []response.OutletProductResponse {
	var returnedResponse []response.OutletProductResponse
	for _, outletPoduct := range outletPoduct {
		singleResponse := response.OutletProductResponse{
			ID:         outletPoduct.ID,
			Name:       outletPoduct.Product.Name,
			OutletID:   outletPoduct.OutletID,
			Price:      outletPoduct.Price,
			OrderPrice: outletPoduct.OrderPrice,
			State:      outletPoduct.State.ToString(),
			CreatedAt:  outletPoduct.CreatedAt,
			UpdatedAt:  outletPoduct.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func OutletProductDetailResponse(outletPoduct models.OutletProduct) response.OutletProductResponse {
	var returnedResponse response.OutletProductResponse
	singleResponse := response.OutletProductResponse{
		ID:         outletPoduct.ID,
		Name:       outletPoduct.Product.Name,
		OutletID:   outletPoduct.OutletID,
		Price:      outletPoduct.Price,
		OrderPrice: outletPoduct.OrderPrice,
		State:      outletPoduct.State.ToString(),
		CreatedAt:  outletPoduct.CreatedAt,
		UpdatedAt:  outletPoduct.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
