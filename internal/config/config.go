package config

import "github.com/caarlos0/env/v9"

type Config struct {
	DatabaseURL        string `env:"DATABASE_URL,notEmpty"`
	ServerPort         string `env:"SERVER_PORT" envDefault:"8080"`
	RedisAddr          string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	MaxConcurrentTasks int    `env:"MAX_CONCURRENT_TASKS" envDefault:"10"`
}

func New() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
