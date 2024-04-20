package service_test

import (
	"context"
	"testing"

	"github.com/dionysia-dev/dionysia/internal/mocks"
	"github.com/dionysia-dev/dionysia/internal/model"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/dionysia-dev/dionysia/internal/task"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPackageStreamSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockClient(ctrl)
	mockStore := mocks.NewMockInputStore(ctrl)
	handler := service.NewNotificationHandler(mockClient, mockStore)

	taskID := uuid.New()
	input := model.Input{ID: taskID}
	expectedTask, _ := task.NewPackageTask(taskID, input)

	expectedInfo := &asynq.TaskInfo{ID: "1", Queue: "default"}

	mockStore.EXPECT().GetInput(gomock.Any(), taskID).Return(input, nil).Times(1)
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
	input := model.Input{ID: taskID}
	expectedTask, _ := task.NewPackageTask(taskID, input)

	mockStore.EXPECT().GetInput(gomock.Any(), taskID).Return(input, nil).Times(1)
	mockClient.EXPECT().Enqueue(expectedTask).Return(nil, assert.AnError).Times(1)

	err := handler.PackageStream(context.TODO(), taskID)

	assert.Error(t, err)
}
