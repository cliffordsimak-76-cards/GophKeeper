package client

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"secret"`
	EnableTLS bool   `env:"ENABLE_TLS" envDefault:"false"`

	ServerHost    string        `env:"GOPHKEEPER_HOST" envDefault:"localhost:9000"`
	ServerTimeout time.Duration `env:"GOPHKEEPER_TIMEOUT" envDefault:"30s"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}

	parseFlags(cfg)

	return cfg, nil
}

// ParseFlags parses the command-line flags from os.Args[1:].
func parseFlags(cfg *Config) {
	flag.BoolVar(&cfg.EnableTLS, "t", cfg.EnableTLS, "enable tls")
	flag.Parse()
}
