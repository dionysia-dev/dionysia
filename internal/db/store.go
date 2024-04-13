package db

import (
	"context"

	"github.com/google/uuid"
)

type InputStore interface {
	CreateInput(context.Context, CreateInputParams) (Input, error)
	GetInput(context.Context, uuid.UUID) (Input, error)
	DeleteInput(context.Context, uuid.UUID) error
	CreateAudioProfile(context.Context, CreateAudioProfileParams) error
	CreateVideoProfile(context.Context, CreateVideoProfileParams) error
	ExecuteTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
