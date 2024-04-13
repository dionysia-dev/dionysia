package service_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/learn-video/dionysia/internal/mocks"
	"github.com/learn-video/dionysia/internal/service"
	"github.com/learn-video/dionysia/internal/task"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPackageStreamSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockClient(ctrl)
	handler := service.NewNotificationHandler(mockClient)

	taskID := uuid.New()
	expectedTask, _ := task.NewPackageTask(taskID)

	expectedInfo := &asynq.TaskInfo{ID: "1", Queue: "default"}

	mockClient.EXPECT().Enqueue(expectedTask).Return(expectedInfo, nil).Times(1)

	err := handler.PackageStream(taskID)

	assert.NoError(t, err)
}

func TestPackageStreamEnqueueFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockClient(ctrl)
	handler := service.NewNotificationHandler(mockClient)

	taskID := uuid.New()
	expectedTask, _ := task.NewPackageTask(taskID)

	mockClient.EXPECT().Enqueue(expectedTask).Return(nil, assert.AnError).Times(1)

	err := handler.PackageStream(taskID)

	assert.Error(t, err)
}
