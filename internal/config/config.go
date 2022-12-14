package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct{}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}
}
