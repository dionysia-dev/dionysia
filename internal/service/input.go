package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/learn-video/streaming-platform/internal/db"
	"github.com/learn-video/streaming-platform/internal/model"
)

type InputStore interface {
	CreateInput(context.Context, db.CreateInputParams) (db.Input, error)
}

type InputHandler interface {
	CreateInput(context.Context, *model.Input) (model.Input, error)
}

type inputHandler struct {
	store InputStore
}

func NewInputHandler(store InputStore) InputHandler {
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
