package task

import (
	"context"
	"fmt"
	"log"

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

func NewPackageTask(uuid uuid.UUID) (*asynq.Task, error) {
	payload, err := json.Marshal(StreamPayload{ID: uuid, Address: "rtmp://localhost:1935"})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeStreamPackage, payload), nil
}

func HandleStreamPackageTask(ctx context.Context, t *asynq.Task) error {
	var p StreamPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal json: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Packaging stream %s", p.ID.String())

	cmd := NewGPACCommand(p.ID.String(), p.Address, "/tmp")
	if err := cmd.Execute(); err != nil {
		log.Printf("Failed to execute command: %v", err)
		return fmt.Errorf("failed to execute command: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Stream %s command executed successfully\n", p.ID.String())

	return nil
}
