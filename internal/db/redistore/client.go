package redistore

import (
	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})
}
