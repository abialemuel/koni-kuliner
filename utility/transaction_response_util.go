package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func TransactionResponse(transaction []models.Transaction) []response.TransactionResponse {
	var returnedResponse []response.TransactionResponse
	for _, transaction := range transaction {
		singleResponse := response.TransactionResponse{
			ID: transaction.ID,
			Customer: response.DetailCustomerResponse{
				ID:      transaction.Customer.ID,
				Name:    transaction.Customer.Name,
				Address: transaction.Customer.Address,
				Phone:   transaction.Customer.Phone,
			},
			Amount:    transaction.Amount,
			State:     transaction.State.ToString(),
			Delivery:  transaction.Delivery.ToString(),
			PoDate:    transaction.PoDate,
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func TransactionDetailResponse(transaction models.Transaction) response.TransactionResponse {
	var returnedResponse response.TransactionResponse
	singleResponse := response.TransactionResponse{
		ID: transaction.ID,
		Customer: response.DetailCustomerResponse{
			ID:      transaction.Customer.ID,
			Name:    transaction.Customer.Name,
			Address: transaction.Customer.Address,
			Phone:   transaction.Customer.Phone,
		},
		Amount:    transaction.Amount,
		State:     transaction.State.ToString(),
		Delivery:  transaction.Delivery.ToString(),
		PoDate:    transaction.PoDate,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}
