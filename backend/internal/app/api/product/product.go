package product

import (
	"context"

	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

type Service interface {
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetWebsiteByAlias(ctx context.Context, alias string) (*model.Website, error)
	GetActiveProductsByAlias(ctx context.Context, alias string) (model.ProductList, error)
}

type Product struct {
	srv Service
}

func NewAPI(srv Service) *Product {
	return &Product{
		srv: srv,
	}
}
