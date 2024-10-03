//go:build migrate

package app

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultAttempts = 20
	defaultTimeout  = time.Second
)

func init() {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) {
		log.Fatal("migrate: environment variable not declared: PG_URL")
	}

	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)

		if err == nil {
			break
		}

		log.Infof("migrate: pgdb is trying to connect, attempts: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("migrate: pgdb connect error: %v", err)
	}

	err = m.Up()
	defer func() { _, _ = m.Close() }()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: up error: %v", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info("migrate: no change")
		return
	}

	log.Infof("migrate: up success")

}
