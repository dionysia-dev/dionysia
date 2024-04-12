package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
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
	expectedID := uuid.New()

	mockStore.EXPECT().CreateInput(ctx, gomock.Any()).Return(db.Input{ID: expectedID, Name: "big buck bunny", Format: "rtmp"}, nil).Times(1)

	result, err := handler.CreateInput(ctx, &input)

	assert.NoError(t, err)
	assert.Equal(t, "big buck bunny", result.Name)
	assert.Equal(t, "rtmp", result.Format)
	assert.Equal(t, expectedID, result.ID)
}

func TestGetInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewInputHandler(mockStore)

	ctx := context.Background()
	expectedID := uuid.New()

	mockStore.EXPECT().GetInput(ctx, expectedID).Return(db.Input{ID: expectedID, Name: "big buck bunny", Format: "rtmp"}, nil).Times(1)

	result, err := handler.GetInput(ctx, expectedID)

	assert.NoError(t, err)
	assert.Equal(t, "big buck bunny", result.Name)
	assert.Equal(t, "rtmp", result.Format)
	assert.Equal(t, expectedID, result.ID)
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
