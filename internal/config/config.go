package config

import "github.com/caarlos0/env/v9"

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,notEmpty"`
	ServerPort  string `env:"SERVER_PORT" envDefault:"8080"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
