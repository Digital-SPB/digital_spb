package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/greenblat17/digital_spb/config"
	"github.com/greenblat17/digital_spb/internal/handlers"
	"github.com/greenblat17/digital_spb/internal/repo"
	"github.com/greenblat17/digital_spb/internal/service"
	data "github.com/greenblat17/digital_spb/pkg/data"
	"github.com/greenblat17/digital_spb/pkg/httpserver"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Database
	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %w", err))
	}
	defer pg.Close()

	// Repositories
	log.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	// Services dependencies
	log.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)

	_ = services

	// Scan Data
	count, err := repositories.EducationalDirection.CountEducationalDirection(context.Background())
	if err != nil {
		log.Fatal("error scanning count educational direction")
	}
	log.Info("count: ", count)
	if count == 0 {
		log.Info("Initializing data...")
		s := data.ScanEducationalDirection()
		for _, v := range s {
			repositories.EducationalDirection.CreateEducationalDirection(context.Background(), v)
		}
	}
	_ = data.ScanEducationalDirection()

	// Handler
	log.Info("Initializing handlers...")
	handler := handlers.NewHandler(services)

	// HTTP server
	log.Info("Starting HTTP server...")
	log.Debugf("Starting port: %s", cfg.HTTP.Port)
	httpServer := httpserver.New(handler.InitRoutes(), httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal:", s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Gracefully shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
