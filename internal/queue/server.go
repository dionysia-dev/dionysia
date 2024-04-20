package queue

import (
	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/hibiken/asynq"
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
