package queue

import (
	"github.com/hibiken/asynq"
	"github.com/learn-video/dionysia/internal/config"
)

type Client interface {
	Enqueue(*asynq.Task, ...asynq.Option) (*asynq.TaskInfo, error)
}

func NewClient(cfg *config.Config) Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisAddr})
}
