package service

import (
	"context"
	"log"

	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/dionysia-dev/dionysia/internal/db"
	"github.com/dionysia-dev/dionysia/internal/queue"
	"github.com/google/uuid"
)

type NotificationHandler interface {
	PackageStream(context.Context, uuid.UUID) error
}

type notificationHandler struct {
	client queue.Client
	store  db.InputStore
	cfg    *config.Config
}

func NewNotificationHandler(c queue.Client, store db.InputStore, cfg *config.Config) NotificationHandler {
	return &notificationHandler{
		client: c,
		store:  store,
		cfg:    cfg,
	}
}

func (h *notificationHandler) PackageStream(ctx context.Context, id uuid.UUID) error {
	in, err := h.store.GetInput(ctx, id)
	if err != nil {
		return err
	}

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

	t, err := NewPackageTask(id, Input{
		ID:            in.ID,
		Name:          in.Name,
		Format:        in.Format,
		VideoProfiles: videoProfiles,
		AudioProfiles: audioProfiles,
	}, h.cfg)
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
