package main

import (
	"context"

	"github.com/Nikita-Kolbin/Maestro/internal/pkg/app"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
)

// @title           Maestro
// @version         1.0

// @host      localhost:8082
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token

func main() {
	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		logger.Error(ctx, "run service failed", "err", err)
	}
}
