package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const (
	TypeStreamPackage = "stream:package"
)

type StreamPayload struct {
	ID      uuid.UUID `json:"id"`
	Address string    `json:"address"`
}

type Notifier interface {
	PackageStream(context.Context)
}

func NewPackageTask(uuid uuid.UUID) (*asynq.Task, error) {
	payload, err := json.Marshal(StreamPayload{ID: uuid, Address: "rtmp://localhost:1935"})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeStreamPackage, payload), nil
}
