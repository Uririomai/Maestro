package customer

import (
	"context"

	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

type Service interface {
	GetJWTSecret() string
	CreateCustomer(ctx context.Context, alias, email, password string) (int, error)
	GetCustomerIdByEmailPassword(ctx context.Context, alias, email, password string) (*model.Customer, error)
}

type Customer struct {
	srv Service
}

func NewAPI(srv Service) *Customer {
	return &Customer{
		srv: srv,
	}
}
