package service

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (s *Service) CreateWebsite(ctx context.Context, alias string, adminId int) (*model.Website, error) {
	haveSite, err := s.repo.AdminHaveWebsite(ctx, adminId)
	if err != nil {
		return nil, err
	}
	if haveSite {
		return nil, model.ErrAdminHaveWebsite
	}

	return s.repo.CreateWebsite(ctx, alias, adminId)
}

func (s *Service) GetWebsiteByAlias(ctx context.Context, alias string) (*model.Website, error) {
	return s.repo.GetWebsiteByAlias(ctx, alias)
}

func (s *Service) GetWebsiteByAdminId(ctx context.Context, adminId int) (*model.Website, error) {
	return s.repo.GetWebsiteByAdminId(ctx, adminId)
}
