package service

import (
	"context"

	"github.com/google/uuid"
)

type OriginStore interface {
	Update(context.Context, Origin) error
	Get(context.Context, uuid.UUID) (Origin, error)
}

type Origin struct {
	ID      uuid.UUID
	Address string
}

type OriginHandler interface {
	Update(context.Context, Origin) error
	Get(context.Context, uuid.UUID) (Origin, error)
}

type originHandler struct {
	store OriginStore
}

func NewOriginHandler(store OriginStore) OriginHandler {
	return &originHandler{
		store: store,
	}
}

func (h *originHandler) Update(ctx context.Context, origin Origin) error {
	return h.store.Update(ctx, origin)
}

func (h *originHandler) Get(ctx context.Context, id uuid.UUID) (Origin, error) {
	return h.store.Get(ctx, id)
}
