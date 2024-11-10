package service

import (
	"context"

	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

type repository interface {
	CreateAdmin(ctx context.Context, email, password string) (int, error)
	GetAdminByEmailPassword(ctx context.Context, email, passwordHash string) (*model.Admin, error)

	CreateWebsite(ctx context.Context, alias string, adminId int) (*model.Website, error)
	GetWebsiteByAlias(ctx context.Context, alias string) (*model.Website, error)
	GetWebsiteByAdminId(ctx context.Context, adminId int) (*model.Website, error)
	AdminHaveWebsite(ctx context.Context, adminId int) (bool, error)

	CreateCustomer(ctx context.Context, alias, email, passwordHash string) (int, error)
	GetCustomerByEmailPassword(ctx context.Context, alias, email, passwordHash string) (*model.Customer, error)

	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetActiveProductsByAlias(ctx context.Context, alias string) (model.ProductList, error)
}

type Service struct {
	jwtSecret string
	repo      repository
}

func New(repo repository, jwtSecret string) *Service {
	return &Service{
		jwtSecret: jwtSecret,
		repo:      repo,
	}
}
