package service

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (s *Service) CreateOrder(ctx context.Context, customerId int, comment string) (*model.Order, error) {
	id, err := s.repo.CreateOrder(ctx, customerId, comment)
	if err != nil {
		return nil, err
	}

	return s.repo.GetOrderById(ctx, id)
}

func (s *Service) GetOrdersByCustomerId(ctx context.Context, customerId int) ([]*model.Order, error) {
	orderIds, err := s.repo.GetOrderIdsByCustomerId(ctx, customerId)
	if err != nil {
		return nil, err
	}

	orders := make([]*model.Order, 0, len(orderIds))
	for _, id := range orderIds {
		order, err := s.repo.GetOrderById(ctx, id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
