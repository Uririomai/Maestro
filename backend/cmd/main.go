package main

import (
	"context"
	
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/app"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
)

func main() {
	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		logger.Error(ctx, "run service failed", "err", err)
	}
}
