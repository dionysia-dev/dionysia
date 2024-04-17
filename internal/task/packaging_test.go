package task_test

import (
	"context"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/learn-video/dionysia/internal/task"
	"github.com/stretchr/testify/assert"
)

func TestHandleStreamPackageTaskFailUnmarshal(t *testing.T) {
	badPayload := []byte("{bad json")
	asynqTask := asynq.NewTask(task.TypeStreamPackage, badPayload)

	err := task.HandleStreamPackageTask(context.Background(), asynqTask)

	assert.Error(t, err)
}
