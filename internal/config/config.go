package config

import (
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	DatabaseURL        string        `env:"DATABASE_URL,notEmpty"`
	APIPort            string        `env:"API_PORT" envDefault:"8080"`
	RedisAddr          string        `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	MaxConcurrentTasks int           `env:"MAX_CONCURRENT_TASKS" envDefault:"10"`
	ReadHeaderTimeout  time.Duration `env:"READ_HEADER_TIMEOUT" envDefault:"2"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
