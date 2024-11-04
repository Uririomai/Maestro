package website

import "context"

type Service interface {
	CreateWebsite(ctx context.Context, alias string, adminId int) error
}

type Website struct {
	srv Service
}

func NewAPI(srv Service) *Website {
	return &Website{
		srv: srv,
	}
}
