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
	GetVideoProfiles(context.Context, uuid.UUID) ([]VideoProfile, error)
	GetAudioProfiles(context.Context, uuid.UUID) ([]AudioProfile, error)
	ExecuteTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
