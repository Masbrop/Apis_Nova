package bootstrap

import (
	"fmt"
	"net/http"
	"time"

	"apis_nova/internal/domain/status"
	handlers "apis_nova/internal/handlers/http"
	"apis_nova/internal/infrastructure/config"
	postgresdb "apis_nova/internal/infrastructure/database/postgres"
	statusrepo "apis_nova/internal/infrastructure/repositories/status"
)

type Application struct {
	Config  config.Config
	Server  *http.Server
	cleanup func() error
}

func NewApplication() (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	repository, cleanup, err := buildStatusRepository(cfg)
	if err != nil {
		return nil, err
	}

	statusService := status.NewService(cfg.AppName, cfg.AppEnv, repository)
	statusHandler := handlers.NewStatusHandler(statusService)

	server := &http.Server{
		Addr:              cfg.HTTPAddress(),
		Handler:           handlers.NewRouter(statusHandler),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return &Application{
		Config:  cfg,
		Server:  server,
		cleanup: cleanup,
	}, nil
}

func (a *Application) Close() error {
	if a.cleanup == nil {
		return nil
	}

	return a.cleanup()
}

func buildStatusRepository(cfg config.Config) (status.Repository, func() error, error) {
	if !cfg.Database.Enabled {
		return statusrepo.NewNoopRepository(), nil, nil
	}

	db, err := postgresdb.Open(cfg.Database)
	if err != nil {
		return nil, nil, fmt.Errorf("open postgres connection: %w", err)
	}

	return statusrepo.NewPostgresRepository(db), db.Close, nil
}
