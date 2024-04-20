package queue

import (
	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/hibiken/asynq"
)

type Client interface {
	Enqueue(*asynq.Task, ...asynq.Option) (*asynq.TaskInfo, error)
}

func NewClient(cfg *config.Config) Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisAddr})
}
