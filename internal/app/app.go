package app

import (
	"github.com/magmaheat/music-library/configs"
	log "github.com/sirupsen/logrus"
)

func Run() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	setupLogger(cfg.Log.Level)

}
