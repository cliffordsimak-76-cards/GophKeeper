package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	GRPCAddr string `env:"GRPC_ADDR" envDefault:":9000"`

	PgDSN          string        `env:"PG_DSN" envDefault:"postgres://***:***@localhost:49153/gophkeeper"`
	PgMaxOpenConn  int           `env:"PG_MAX_OPEN_CONN" envDefault:"5"`
	PgIdleConn     int           `env:"PG_MAX_IDLE_CONN" envDefault:"5"`
	PgPingInterval time.Duration `env:"PG_PING_INTERVAL" envDefault:"10m"`

	SecretKey     string        `env:"SECRET_KEY" envDefault:"secret"`
	TokenDuration time.Duration `env:"TOKEN_DURATION" envDefault:"15m"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}

	return cfg, nil
}
