package service

import (
	"context"
	"io"

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

type objectStorage interface {
	PutObject(ctx context.Context, reader io.Reader, size int64, bucketName, contentType string) (_ string, err error)
	GetObject(ctx context.Context, objectId, bucketName string) (io.Reader, string, error)
}

type Service struct {
	jwtSecret string
	repo      repository
	storage   objectStorage
}

func New(repo repository, storage objectStorage, jwtSecret string) *Service {
	return &Service{
		jwtSecret: jwtSecret,
		repo:      repo,
		storage:   storage,
	}
}
