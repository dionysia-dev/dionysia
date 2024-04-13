package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/mock/gomock"

	"github.com/learn-video/dionysia/internal/db"
	"github.com/learn-video/dionysia/internal/mocks"
	"github.com/learn-video/dionysia/internal/model"
	"github.com/learn-video/dionysia/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewInputHandler(mockStore)

	ctx := context.Background()
	input := model.Input{Name: "big buck bunny", Format: "rtmp"}

	mockStore.EXPECT().ExecuteTransaction(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
		return fn(ctx)
	}).Times(1)

	mockStore.EXPECT().CreateInput(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, params db.CreateInputParams) (db.Input, error) {
		assert.Equal(t, "big buck bunny", params.Name)
		assert.Equal(t, "rtmp", params.Format)
		return db.Input(params), nil
	}).Times(1)

	result, err := handler.CreateInput(ctx, &input)

	assert.NoError(t, err)
	assert.Equal(t, "big buck bunny", result.Name)
	assert.Equal(t, "rtmp", result.Format)
}

func TestGetInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewInputHandler(mockStore)

	ctx := context.Background()
	expectedID := uuid.New()

	// Mock the call to GetInput
	mockStore.EXPECT().GetInput(ctx, expectedID).Return(db.Input{
		ID:     expectedID,
		Name:   "big buck bunny",
		Format: "rtmp",
	}, nil).Times(1)

	mockStore.EXPECT().GetVideoProfiles(ctx, expectedID).Return([]db.VideoProfile{
		{Codec: "h264", Bitrate: pgtype.Int4{Int32: 1000}},
	}, nil).Times(1)

	mockStore.EXPECT().GetAudioProfiles(ctx, expectedID).Return([]db.AudioProfile{
		{Codec: "aac", Rate: pgtype.Int4{Int32: 128}},
	}, nil).Times(1)

	result, err := handler.GetInput(ctx, expectedID)

	assert.NoError(t, err)
	assert.Equal(t, "big buck bunny", result.Name)
	assert.Equal(t, "rtmp", result.Format)
	assert.Equal(t, expectedID, result.ID)
	assert.Len(t, result.Video, 1)
	assert.Len(t, result.Audio, 1)
	assert.Equal(t, "h264", result.Video[0].Codec)
	assert.Equal(t, 1000, result.Video[0].Bitrate)
	assert.Equal(t, "aac", result.Audio[0].Codec)
	assert.Equal(t, 128, result.Audio[0].Rate)
}

func TestDeleteInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewInputHandler(mockStore)

	ctx := context.Background()
	inputID := uuid.New()

	mockStore.EXPECT().DeleteInput(ctx, inputID).Return(nil).Times(1)

	err := handler.DeleteInput(ctx, inputID)

	assert.NoError(t, err)
}
