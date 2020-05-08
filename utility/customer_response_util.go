package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func CustomerResponse(customer []models.Customer) []response.CustomerResponse {
	var returnedResponse []response.CustomerResponse
	for _, customer := range customer {
		singleResponse := response.CustomerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			Address:   customer.Address,
			Phone:     customer.Phone,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func CustomerDetailResponse(customer models.Customer) response.CustomerResponse {
	var returnedResponse response.CustomerResponse
	singleResponse := response.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Address:   customer.Address,
		Phone:     customer.Phone,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
