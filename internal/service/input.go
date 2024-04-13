package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	var input db.Input

	err := handler.store.ExecuteTransaction(ctx, func(txCtx context.Context) error {
		var err error
		input, err = handler.store.CreateInput(txCtx, db.CreateInputParams{
			ID:     uuid.New(),
			Name:   in.Name,
			Format: in.Format,
		})

		if err != nil {
			return err
		}

		for _, a := range in.Audio {
			if err := handler.store.CreateAudioProfile(ctx, db.CreateAudioProfileParams{
				InputID: input.ID,
				Rate:    pgtype.Int4{Int32: int32(a.Rate), Valid: true},
				Codec:   a.Codec,
			}); err != nil {
				return err
			}
		}

		for _, v := range in.Video {
			if err := handler.store.CreateVideoProfile(ctx, db.CreateVideoProfileParams{
				InputID:        input.ID,
				Width:          pgtype.Int4{Int32: int32(v.Width), Valid: true},
				Height:         pgtype.Int4{Int32: int32(v.Height), Valid: true},
				Codec:          v.Codec,
				MaxKeyInterval: pgtype.Int4{Int32: int32(v.MaxKeyInterval), Valid: true},
				Framerate:      pgtype.Int4{Int32: int32(v.Framerate), Valid: true},
				Bitrate:        pgtype.Int4{Int32: int32(v.Bitrate), Valid: true},
			}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return model.Input{}, err
	}

	return model.Input{
		ID:     input.ID,
		Name:   input.Name,
		Format: input.Format,
		Audio:  in.Audio,
		Video:  in.Video,
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
