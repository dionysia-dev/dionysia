package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

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

	mockStore.EXPECT().CreateInput(ctx, &input).Return(nil).Times(1)

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
	expectedInput := model.Input{
		ID:            expectedID,
		Name:          "big buck bunny",
		Format:        "rtmp",
		VideoProfiles: []model.VideoProfile{{Codec: "h264", Bitrate: 1000}},
		AudioProfiles: []model.AudioProfile{{Codec: "aac", Rate: 128}},
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
	assert.Equal(t, 128, result.AudioProfiles[0].Rate)
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
