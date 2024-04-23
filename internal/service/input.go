package service

import (
	"context"
	"errors"

	"github.com/dionysia-dev/dionysia/internal/db"
	"github.com/dionysia-dev/dionysia/internal/db/model"
	"github.com/google/uuid"
)

var (
	ErrFailedAuth = errors.New("input not found")
)

type Input struct {
	ID            uuid.UUID
	Name          string
	Format        string
	AudioProfiles []AudioProfile
	VideoProfiles []VideoProfile
}

type AudioProfile struct {
	InputID uuid.UUID
	Codec   string
	Bitrate int
}

type VideoProfile struct {
	InputID        uuid.UUID
	Codec          string
	Bitrate        int
	MaxKeyInterval int
	Framerate      int
	Width          int
	Height         int
}

type IngestAuth struct {
	Path   uuid.UUID
	Action string
}

type InputHandler interface {
	CreateInput(context.Context, *Input) (Input, error)
	GetInput(context.Context, uuid.UUID) (Input, error)
	DeleteInput(context.Context, uuid.UUID) error
	Authenticate(context.Context, IngestAuth) error
}

type inputHandler struct {
	store db.InputStore
}

func NewInputHandler(store db.InputStore) InputHandler {
	return &inputHandler{
		store: store,
	}
}

func (handler *inputHandler) CreateInput(ctx context.Context, in *Input) (Input, error) {
	in.ID = uuid.New()

	audioProfiles := make([]model.AudioProfile, len(in.AudioProfiles))
	for i, audio := range in.AudioProfiles {
		audioProfiles[i] = model.AudioProfile{
			InputID: in.ID,
			Codec:   audio.Codec,
			Bitrate: audio.Bitrate,
		}
	}

	videoProfiles := make([]model.VideoProfile, len(in.VideoProfiles))
	for i, video := range in.VideoProfiles {
		videoProfiles[i] = model.VideoProfile{
			InputID:        in.ID,
			Codec:          video.Codec,
			Bitrate:        video.Bitrate,
			MaxKeyInterval: video.MaxKeyInterval,
			Framerate:      video.Framerate,
			Width:          video.Width,
			Height:         video.Height,
		}
	}

	err := handler.store.CreateInput(ctx, &model.Input{
		ID:            in.ID,
		Name:          in.Name,
		Format:        in.Format,
		AudioProfiles: audioProfiles,
		VideoProfiles: videoProfiles,
	})

	return *in, err
}

func (handler *inputHandler) GetInput(ctx context.Context, id uuid.UUID) (Input, error) {
	in, err := handler.store.GetInput(ctx, id)

	audioProfiles := make([]AudioProfile, len(in.AudioProfiles))
	for i, audio := range in.AudioProfiles { //nolint: gocritic // Can't use pointers in models
		audioProfiles[i] = AudioProfile{
			InputID: in.ID,
			Codec:   audio.Codec,
			Bitrate: audio.Bitrate,
		}
	}

	videoProfiles := make([]VideoProfile, len(in.VideoProfiles))
	for i, video := range in.VideoProfiles { //nolint: gocritic // Can't use pointers in models
		videoProfiles[i] = VideoProfile{
			InputID:        in.ID,
			Codec:          video.Codec,
			Bitrate:        video.Bitrate,
			MaxKeyInterval: video.MaxKeyInterval,
			Framerate:      video.Framerate,
			Width:          video.Width,
			Height:         video.Height,
		}
	}

	return Input{
		ID:            in.ID,
		Name:          in.Name,
		Format:        in.Format,
		AudioProfiles: audioProfiles,
		VideoProfiles: videoProfiles,
	}, err
}

func (handler *inputHandler) DeleteInput(ctx context.Context, id uuid.UUID) error {
	return handler.store.DeleteInput(ctx, id)
}

func (handler *inputHandler) Authenticate(ctx context.Context, authData IngestAuth) error {
	_, err := handler.store.GetInput(ctx, authData.Path)
	if errors.Is(err, db.ErrNotFound) {
		return ErrFailedAuth
	}

	return nil
}
