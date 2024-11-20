package service

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (s *Service) AddProductToCart(ctx context.Context, cartId, productId, count int) (*model.Cart, error) {
	err := s.repo.UpsertCartItem(ctx, cartId, productId, count)
	if err != nil {
		return nil, err
	}

	return s.repo.GetCart(ctx, cartId)
}
