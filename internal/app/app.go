package app

import (
	"github.com/labstack/echo/v4"
	"github.com/magmaheat/music-info/configs"
	"github.com/magmaheat/music-info/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

func Run() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	log.Infof("setup log level: %s", cfg.Level)
	setupLogger(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(20))
	if err != nil {
		log.Fatalf("postgres error: %v", err)
	}

	_ = pg

	handler := echo.New()

	log.Infof("start server on ports: %s", cfg.Port)
	if err := handler.Start(":" + cfg.Port); err != nil {
		log.Fatalf("server stoped ...")
	}
}
