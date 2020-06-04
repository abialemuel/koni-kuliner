package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func SellerResponse(seller []models.Seller) []response.SellerResponse {
	var returnedResponse []response.SellerResponse
	for _, seller := range seller {
		singleResponse := response.SellerResponse{
			ID:        seller.ID,
			Name:      seller.Name,
			Username:  seller.Username,
			Address:   seller.Address,
			Phone:     seller.Phone,
			CreatedAt: seller.CreatedAt,
			UpdatedAt: seller.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func SellerDetailResponse(seller models.Seller) response.SellerResponse {
	var returnedResponse response.SellerResponse
	singleResponse := response.SellerResponse{
		ID:        seller.ID,
		Name:      seller.Name,
		Username:  seller.Username,
		Address:   seller.Address,
		Phone:     seller.Phone,
		CreatedAt: seller.CreatedAt,
		UpdatedAt: seller.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
