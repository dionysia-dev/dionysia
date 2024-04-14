package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/learn-video/dionysia/internal/db"
	"github.com/learn-video/dionysia/internal/queue"
	"github.com/learn-video/dionysia/internal/task"
)

type NotificationHandler interface {
	PackageStream(context.Context, uuid.UUID) error
}

type notificationHandler struct {
	client queue.Client
	store  db.InputStore
}

func NewNotificationHandler(c queue.Client, store db.InputStore) NotificationHandler {
	return &notificationHandler{
		client: c,
		store:  store,
	}
}

func (h *notificationHandler) PackageStream(ctx context.Context, id uuid.UUID) error {
	input, err := h.store.GetInput(ctx, id)
	if err != nil {
		return err
	}

	t, err := task.NewPackageTask(id, input)
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
