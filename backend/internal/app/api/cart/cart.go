package cart

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

type Service interface {
	GetProductById(ctx context.Context, id int) (*model.Product, error)
	AddProductToCart(ctx context.Context, cartId, productId, count int) (*model.Cart, error)
}

type Cart struct {
	srv Service
}

func NewAPI(srv Service) *Cart {
	return &Cart{
		srv: srv,
	}
}
