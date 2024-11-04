package service

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (s *Service) CreateCustomer(ctx context.Context, alias, email, password string) (int, error) {
	if !validEmail(email) {
		return 0, model.ErrInvalidEmail
	}

	hash := generatePasswordHash(password)

	return s.repo.CreateCustomer(ctx, alias, email, hash)
}

func (s *Service) GetCustomerIdByEmailPassword(ctx context.Context, alias, email, password string) (*model.Customer, error) {
	if !validEmail(email) {
		return nil, model.ErrInvalidEmail
	}

	hash := generatePasswordHash(password)

	return s.repo.GetCustomerByEmailPassword(ctx, alias, email, hash)
}
