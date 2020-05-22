package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func TransactionResponse(transaction []models.Transaction) []response.TransactionResponse {
	var returnedResponse []response.TransactionResponse
	for _, transaction := range transaction {
		singleResponse := response.TransactionResponse{
			ID:        transaction.ID,
			Customer:  getDetailCustomerResponse(transaction),
			CartItems: getDetailCartItemResponse(transaction),
			Amount:    transaction.Amount,
			State:     transaction.State.ToString(),
			Delivery:  transaction.Delivery.ToString(),
			Note:      transaction.Note,
			PoDate:    transaction.PoDate,
			CreatedAt: transaction.CreatedAt,
			UpdatedAt: transaction.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

func TransactionDetailResponse(transaction models.Transaction) response.TransactionDetailResponse {
	var returnedResponse response.TransactionDetailResponse
	singleResponse := response.TransactionDetailResponse{
		ID:        transaction.ID,
		Customer:  getDetailCustomerResponse(transaction),
		CartItems: getDetailCartItemResponse(transaction),
		Amount:    transaction.Amount,
		State:     transaction.State.ToString(),
		Delivery:  transaction.Delivery.ToString(),
		Note:      transaction.Note,
		Feedback:  transaction.Feedback,
		PoDate:    transaction.PoDate,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}
	returnedResponse = singleResponse
	return returnedResponse
}

// private func

func getDetailCartItemResponse(transaction models.Transaction) []response.DetailCartItemResponse {
	var cartItems []response.DetailCartItemResponse
	for _, cartItem := range transaction.CartItems {
		singleCart := response.DetailCartItemResponse{
			ID: cartItem.ID,
			Product: response.DetailProductResponse{
				ID:   cartItem.OutletProduct.Product.ID,
				Name: cartItem.OutletProduct.Product.Name,
			},
			Quantity: cartItem.Quantity,
			Price:    cartItem.Price,
		}
		cartItems = append(cartItems, singleCart)
	}
	return cartItems
}

func getDetailCustomerResponse(transaction models.Transaction) response.DetailCustomerResponse {
	var customer response.DetailCustomerResponse
	customer = response.DetailCustomerResponse{
		ID:      transaction.Customer.ID,
		Name:    transaction.Customer.Name,
		Address: transaction.Customer.Address,
		Phone:   transaction.Customer.Phone,
	}
	return customer
}
