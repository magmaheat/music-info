package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/magmaheat/music-info/configs"
	v1 "github.com/magmaheat/music-info/internal/controller/http/v1"
	"github.com/magmaheat/music-info/internal/repo"
	"github.com/magmaheat/music-info/internal/service"
	"github.com/magmaheat/music-info/pkg/httpserver"
	"github.com/magmaheat/music-info/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title Music Info
// @version 0.0.1
// @description ... //TODO add description

// @contact.name George Epishev
// @contact.email epishcom@gmail.com

// @host localhost:8090
// @BasePath /

func Run() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	setupLogger(cfg.Log.Level)

	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(20))
	if err != nil {
		log.Fatalf("postgres error: %v", err)
	}
	defer pg.Close()

	log.Info("Initializing repositories...")

	repositories := repo.New(pg)

	log.Info("Initializing services...")
	services := service.NewServices(repositories)

	log.Info("Initializing handlers and router")
	handler := echo.New()
	v1.NewRouter(handler, services)

	log.Infof("Starting server, port: %s", cfg.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %v", err))
	}

	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServerShutdown: %v", err))
	}
}
