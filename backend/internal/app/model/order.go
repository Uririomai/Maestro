package model

import "time"

type Order struct {
	Id         int       `db:"id"`
	CustomerId int       `db:"customer_id"`
	TotalSum   int       `db:"total_sum"`
	Date       time.Time `db:"date_time"`
	Status     int       `db:"status"`
	Comment    string    `db:"comment"`
	Items      []*OrderItem
}

type OrderItem struct {
	Id      int `db:"id"`
	OrderId int `db:"order_id"`
	Count   int `db:"count"`
	Product
}

type OrderDTO struct {
	Id         int             `json:"id"`
	CustomerId int             `json:"customer_id"`
	TotalSum   int             `json:"total_sum"`
	Date       string          `json:"date_time"`
	Status     int             `json:"status"`
	Comment    string          `json:"comment"`
	Items      []*OrderItemDTO `json:"items"`
}

type OrderItemDTO struct {
	Product *ProductDTO `json:"product"`
	Count   int         `json:"count"`
}

type MakeOrderRequest struct {
	Comment string `json:"comment"`
}

type GetMyOrdersResponse struct {
	Orders []*OrderDTO `json:"orders"`
}

func FromOrderToDTO(order *Order) *OrderDTO {
	resp := &OrderDTO{
		Id:         order.Id,
		CustomerId: order.CustomerId,
		TotalSum:   order.TotalSum,
		Date:       order.Date.Format(time.DateTime),
		Status:     order.Status,
		Comment:    order.Comment,
		Items:      make([]*OrderItemDTO, 0, len(order.Items)),
	}

	for _, item := range order.Items {
		resp.Items = append(resp.Items, &OrderItemDTO{
			Product: FromProductToDTO(&item.Product),
			Count:   item.Count,
		})
	}

	return resp
}

func FromOrdersToDTO(orders []*Order) []*OrderDTO {
	resp := make([]*OrderDTO, 0, len(orders))
	for _, order := range orders {
		resp = append(resp, FromOrderToDTO(order))
	}
	return resp
}
