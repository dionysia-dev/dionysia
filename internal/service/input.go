package service

import (
	"context"
	"errors"

	"github.com/dionysia-dev/dionysia/internal/db"
	"github.com/dionysia-dev/dionysia/internal/model"
	"github.com/google/uuid"
)

var (
	ErrFailedAuth = errors.New("input not found")
)

type InputHandler interface {
	CreateInput(context.Context, *model.Input) (model.Input, error)
	GetInput(context.Context, uuid.UUID) (model.Input, error)
	DeleteInput(context.Context, uuid.UUID) error
	Authenticate(context.Context, model.IngestAuthData) error
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

func (handler *inputHandler) Authenticate(ctx context.Context, authData model.IngestAuthData) error {
	_, err := handler.store.GetInput(ctx, authData.Path)
	if errors.Is(err, db.ErrNotFound) {
		return ErrFailedAuth
	}

	return nil
}
