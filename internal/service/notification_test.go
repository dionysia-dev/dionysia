package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"go.uber.org/mock/gomock"

	"github.com/dionysia-dev/dionysia/internal/db/model" // Make sure this is the correct import path for model
	"github.com/dionysia-dev/dionysia/internal/mocks"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestPackageStreamSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockClient(ctrl)
	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewNotificationHandler(mockClient, mockStore)

	taskID := uuid.New()
	expectedInput := model.Input{ID: taskID, Name: "test"}
	expectedTask, _ := service.NewPackageTask(taskID, service.Input{ID: taskID, Name: "test"})

	expectedInfo := &asynq.TaskInfo{ID: "1", Queue: "default"}

	mockStore.EXPECT().GetInput(gomock.Any(), taskID).Return(expectedInput, nil).Times(1)
	mockClient.EXPECT().Enqueue(expectedTask).Return(expectedInfo, nil).Times(1)

	err := handler.PackageStream(context.TODO(), taskID)

	assert.NoError(t, err)
}

func TestPackageStreamEnqueueFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockClient(ctrl)
	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewNotificationHandler(mockClient, mockStore)

	taskID := uuid.New()
	expectedInput := model.Input{ID: taskID, Name: "test"}
	expectedTask, _ := service.NewPackageTask(taskID, service.Input{ID: taskID, Name: "test"})

	mockStore.EXPECT().GetInput(gomock.Any(), taskID).Return(expectedInput, nil).Times(1)
	mockClient.EXPECT().Enqueue(expectedTask).Return(nil, assert.AnError).Times(1)

	err := handler.PackageStream(context.TODO(), taskID)

	assert.Error(t, err)
}
