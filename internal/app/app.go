package app

import (
	"DobryySoul/project-with-API-interaction/internal/config"
	"DobryySoul/project-with-API-interaction/internal/http/handler"
	"DobryySoul/project-with-API-interaction/internal/http/router"
	"DobryySoul/project-with-API-interaction/internal/services/service"
	"DobryySoul/project-with-API-interaction/internal/storage/repo"
	"DobryySoul/project-with-API-interaction/pkg/logger"
	"DobryySoul/project-with-API-interaction/pkg/postgres"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	l, err := logger.New()
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	pg, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer pg.Close()

	r := gin.Default()
	userRepo := repo.NewUserRepository(pg, l)
	service := service.NewService(userRepo, l)
	handler := handler.NewHandlers(service, cfg, l)

	router.NewRouter(r, handler)

	if err := r.Run(cfg.HTTP.Host + ":" + cfg.HTTP.Port); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}
