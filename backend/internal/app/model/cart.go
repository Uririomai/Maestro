package model

type Cart struct {
	Items []*CartItem
}

type CartItem struct {
	Id     int `db:"id"`
	CartId int `db:"cart_id"`
	Count  int `db:"count"`
	Product
}

type AddProductToCartRequest struct {
	ProductId int `json:"product_id"`
	Count     int `json:"count"`
}

type CartResponse struct {
	Items []*CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	Product *ProductDTO `json:"product"`
	Count   int         `json:"count"`
}

func FromCartToDTO(cart *Cart) *CartResponse {
	resp := &CartResponse{
		Items: make([]*CartItemResponse, 0, len(cart.Items)),
	}

	for _, item := range cart.Items {
		resp.Items = append(resp.Items, &CartItemResponse{
			Product: FromProductToDTO(&item.Product),
			Count:   item.Count,
		})
	}

	return resp
}
