package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/learn-video/dionysia/internal/task"
)

type NotificationHandler interface {
	PackageStream(id uuid.UUID) error
}

type notificationHandler struct {
	client *asynq.Client
}

func NewNotificationHandler(c *asynq.Client) NotificationHandler {
	return &notificationHandler{
		client: c,
	}
}

func (h *notificationHandler) PackageStream(id uuid.UUID) error {
	t, err := task.NewPackageTask(id)
	if err != nil {
		return err
	}

	info, err := h.client.Enqueue(t)
	if err != nil {
		return err
	}

	log.Printf("Package Stream enqueued: %s %s\n", info.ID, info.Queue)

	return nil
}
