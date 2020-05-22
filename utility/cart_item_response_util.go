package utility

import (
	"github.com/koni-kuliner/models"
	"github.com/koni-kuliner/resource/response"
)

func CartItemResponse(cartItem []models.CartItem) []response.CartItemResponse {
	var returnedResponse []response.CartItemResponse
	for _, cartItem := range cartItem {
		singleResponse := response.CartItemResponse{
			ID:              cartItem.ID,
			Product:         getDetailProductResponse(cartItem),
			OutletProductID: cartItem.OutletProductID,
			Price:           cartItem.Price,
			OrderPrice:      cartItem.OrderPrice,
			CreatedAt:       cartItem.CreatedAt,
			UpdatedAt:       cartItem.UpdatedAt,
		}
		returnedResponse = append(returnedResponse, singleResponse)
	}
	return returnedResponse
}

// func CartItemDetailResponse(cartItem models.CartItem) response.CartItemDetailResponse {
// 	var returnedResponse response.CartItemDetailResponse
// 	singleResponse := response.CartItemDetailResponse{
// 		ID:        cartItem.ID,
// 		Customer:  getDetailCustomerResponse(cartItem),
// 		CartItems: getDetailCartItemResponse(cartItem),
// 		Amount:    cartItem.Amount,
// 		State:     cartItem.State.ToString(),
// 		Delivery:  cartItem.Delivery.ToString(),
// 		Note:      cartItem.Note,
// 		Feedback:  cartItem.Feedback,
// 		PoDate:    cartItem.PoDate,
// 		CreatedAt: cartItem.CreatedAt,
// 		UpdatedAt: cartItem.UpdatedAt,
// 	}
// 	returnedResponse = singleResponse
// 	return returnedResponse
// }

// private func

func getDetailProductResponse(cartItem models.CartItem) response.DetailProductResponse {
	var product response.DetailProductResponse
	product = response.DetailProductResponse{
		ID:   cartItem.OutletProduct.Product.ID,
		Name: cartItem.OutletProduct.Product.Name,
	}
	return product
}
