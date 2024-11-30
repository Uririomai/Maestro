package service

import (
	"context"

	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (s *Service) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	return s.repo.CreateProduct(ctx, product)
}

func (s *Service) UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	return s.repo.UpdateProduct(ctx, product)
}

func (s *Service) GetProductById(ctx context.Context, id int) (*model.Product, error) {
	return s.repo.GetProductById(ctx, id)
}

func (s *Service) GetActiveProductsByAlias(ctx context.Context, alias string) (model.ProductList, error) {
	return s.repo.GetActiveProductsByAlias(ctx, alias)
}
