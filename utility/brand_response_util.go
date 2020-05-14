package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func BrandResponse(brand []models.Brand) []response.BrandResponse {
	var returnedResponse []response.BrandResponse
	for _, brand := range brand {
		singleResponse := response.BrandResponse{
			ID:        brand.ID,
			Name:      brand.Name,
			CreatedAt: brand.CreatedAt,
			UpdatedAt: brand.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func BrandDetailResponse(brand models.Brand) response.BrandResponse {
	var returnedResponse response.BrandResponse
	singleResponse := response.BrandResponse{
		ID:        brand.ID,
		Name:      brand.Name,
		CreatedAt: brand.CreatedAt,
		UpdatedAt: brand.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
