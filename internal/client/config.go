package client

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"secret"`

	ServerHost    string        `env:"GOPHKEEPER_HOST" envDefault:"localhost:9000"`
	ServerTimeout time.Duration `env:"GOPHKEEPER_TIMEOUT" envDefault:"30s"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}

	return cfg, nil
}
