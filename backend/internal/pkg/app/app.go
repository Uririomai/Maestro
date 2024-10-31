package app

import (
	"context"
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/api/router"
	"github.com/Nikita-Kolbin/Maestro/internal/app/repository"
	"github.com/Nikita-Kolbin/Maestro/internal/app/service"
)

func Run(ctx context.Context) error {
	repo, err := repository.New()
	if err != nil {
		return err
	}

	srv := service.New(repo)

	r := router.New(srv)

	_ = r

	return errors.New("errr aaaaa")
}
