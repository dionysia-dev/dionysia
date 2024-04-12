package queue

import (
	"github.com/hibiken/asynq"
	"github.com/learn-video/dionysia/internal/config"
)

func NewServer(cfg *config.Config) *asynq.Server {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.RedisAddr},
		asynq.Config{
			Concurrency: cfg.MaxConcurrentTasks,
		},
	)

	return srv
}
