package service_test

import (
	"context"
	"testing"

	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestHandleStreamPackageTaskFailUnmarshal(t *testing.T) {
	badPayload := []byte("{bad json")
	asynqTask := asynq.NewTask(service.TypeStreamPackage, badPayload)

	err := service.HandleStreamPackageTask(context.Background(), asynqTask)

	assert.Error(t, err)
}
