package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func ProductResponse(product []models.Product) []response.ProductResponse {
	var returnedResponse []response.ProductResponse
	for _, product := range product {
		singleResponse := response.ProductResponse{
			ID:   product.ID,
			Name: product.Name,
			Brand: response.DetailBrandResponse{
				ID:   product.Brand.ID,
				Name: product.Brand.Name,
			},
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func ProductDetailResponse(product models.Product) response.ProductResponse {
	var returnedResponse response.ProductResponse
	singleResponse := response.ProductResponse{
		ID:   product.ID,
		Name: product.Name,
		Brand: response.DetailBrandResponse{
			ID:   product.Brand.ID,
			Name: product.Brand.Name,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
