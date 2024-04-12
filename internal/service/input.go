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
	input, err := handler.store.CreateInput(ctx, db.CreateInputParams{
		ID:     uuid.New(),
		Name:   in.Name,
		Format: in.Format,
	})

	if err != nil {
		return model.Input{}, err
	}

	return model.Input{
		ID:     input.ID,
		Name:   input.Name,
		Format: input.Format,
	}, nil
}

func (handler *inputHandler) GetInput(ctx context.Context, id uuid.UUID) (model.Input, error) {
	input, err := handler.store.GetInput(ctx, id)
	if err != nil {
		return model.Input{}, err
	}

	return model.Input{
		ID:     input.ID,
		Name:   input.Name,
		Format: input.Format,
	}, nil
}

func (handler *inputHandler) DeleteInput(ctx context.Context, id uuid.UUID) error {
	return handler.store.DeleteInput(ctx, id)
}
