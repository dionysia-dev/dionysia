package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/learn-video/dionysia/internal/db"
	"github.com/learn-video/dionysia/internal/model"
)

type InputHandler interface {
	CreateInput(context.Context, *model.Input) (model.Input, error)
	GetInput(context.Context, uuid.UUID) (model.Input, error)
	DeleteInput(context.Context, uuid.UUID) error
}

type inputHandler struct {
	store db.InputStore
}

func NewInputHandler(store db.InputStore) InputHandler {
	return &inputHandler{
		store: store,
	}
}

func (handler *inputHandler) CreateInput(ctx context.Context, in *model.Input) (model.Input, error) {
	in.ID = uuid.New()
	err := handler.store.CreateInput(ctx, in)

	return *in, err
}

func (handler *inputHandler) GetInput(ctx context.Context, id uuid.UUID) (model.Input, error) {
	return handler.store.GetInput(ctx, id)
}

func (handler *inputHandler) DeleteInput(ctx context.Context, id uuid.UUID) error {
	return handler.store.DeleteInput(ctx, id)
}
