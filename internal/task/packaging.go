package task

import (
	"context"
	"fmt"
	"log"

	"encoding/json"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/learn-video/dionysia/internal/model"
)

const (
	TypeStreamPackage = "stream:package"
)

type StreamPayload struct {
	ID      uuid.UUID   `json:"id"`
	Input   model.Input `json:"input"`
	Address string      `json:"address"`
}

func NewPackageTask(id uuid.UUID, input model.Input) (*asynq.Task, error) {
	payload, err := json.Marshal(StreamPayload{
		ID:      id,
		Input:   input,
		Address: "rtmp://media-server:1935",
	})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeStreamPackage, payload), nil
}

func HandleStreamPackageTask(_ context.Context, t *asynq.Task) error {
	var p StreamPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal json: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Packaging stream %s", p.ID.String())

	cfg := NewDefaultCommandConfig()
	builder := NewGPACCommandBuilder(cfg)
	cmd := builder.Build(p.ID.String(), p.Address, "/output", p.Input)

	if err := cmd.Execute(); err != nil {
		log.Printf("Failed to execute command: %v", err)

		return fmt.Errorf("failed to execute command: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Stream %s command executed successfully\n", p.ID.String())

	return nil
}
