package app

import (
	"github.com/labstack/echo/v4"
	"github.com/magmaheat/music-info/configs"
	"github.com/magmaheat/music-info/internal/repo"
	"github.com/magmaheat/music-info/internal/service"
	"github.com/magmaheat/music-info/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

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

	_ = pg

	log.Info("Initializing repositories...")

	repositories := repo.New(pg)

	log.Info("Initializing services...")
	services := service.NewServices(repositories)

	log.Info("Initializing handlers and router")
	_ = services
	handler := echo.New()

	log.Infof("start server on ports: %s", cfg.Port)
	if err := handler.Start(":" + cfg.Port); err != nil {
		log.Fatalf("server stoped ...")
	}
}
