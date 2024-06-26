package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	"github.com/dionysia-dev/dionysia/internal/db/model"
	"github.com/dionysia-dev/dionysia/internal/mocks"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewInputHandler(mockStore)

	ctx := context.Background()

	mockStore.EXPECT().CreateInput(ctx, gomock.Any()).DoAndReturn(func(_ context.Context, in *model.Input) *model.Input {
		in.ID = uuid.New()
		return in
	}).Return(nil).Times(1)

	result, err := handler.CreateInput(ctx, &service.Input{Name: "big buck bunny", Format: "rtmp"})

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
	expectedInput := model.Input{
		ID:            expectedID,
		Name:          "big buck bunny",
		Format:        "rtmp",
		VideoProfiles: []model.VideoProfile{{Codec: "h264", Bitrate: 1000}},
		AudioProfiles: []model.AudioProfile{{Codec: "aac", Bitrate: 128}},
	}

	mockStore.EXPECT().GetInput(ctx, expectedID).Return(expectedInput, nil).Times(1)

	result, err := handler.GetInput(ctx, expectedID)

	assert.NoError(t, err)
	assert.Equal(t, "big buck bunny", result.Name)
	assert.Equal(t, "rtmp", result.Format)
	assert.Equal(t, expectedID, result.ID)
	assert.Len(t, result.VideoProfiles, 1)
	assert.Len(t, result.AudioProfiles, 1)
	assert.Equal(t, "h264", result.VideoProfiles[0].Codec)
	assert.Equal(t, 1000, result.VideoProfiles[0].Bitrate)
	assert.Equal(t, "aac", result.AudioProfiles[0].Codec)
	assert.Equal(t, 128, result.AudioProfiles[0].Bitrate)
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
