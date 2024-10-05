package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App
	Log
	HTTP
	PG
}

type App struct {
	Name    string `env-required:"true" env:"APP_NAME"`
	Version string `env-required:"true" env:"APP_VERSION"`
}

type Log struct {
	Level string `env-required:"true" env:"LOG_LEVEL"`
}

type HTTP struct {
	Host string `env-required:"true" env:"HTTP_HOST"`
	Port string `env-required:"true" env:"HTTP_PORT"`
}

type PG struct {
	PoolSize int    `env-required:"true" env:"PG_POOL_SIZE"`
	URL      string `env-required:"true" env:"PG_URL"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error update environments: %v", err)
	}

	return cfg, nil
}
