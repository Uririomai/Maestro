package service

import (
	"context"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
)

func (s *Service) CreateWebsite(ctx context.Context, alias string, adminId int) error {
	haveSite, err := s.repo.AdminHaveWebsite(ctx, adminId)
	if err != nil {
		return err
	}
	if haveSite {
		return model.ErrAdminHaveWebsite
	}

	return s.repo.CreateWebsite(ctx, alias, adminId)
}
