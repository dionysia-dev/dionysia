package service_test

import (
	"context"
	"testing"

	"github.com/dionysia-dev/dionysia/internal/mocks"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOriginHandlerUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockOriginStore(ctrl)
	handler := service.NewOriginHandler(mockStore)

	ctx := context.Background()
	origin := service.Origin{
		ID:      uuid.New(),
		Address: "http://localhost:9999",
	}

	mockStore.EXPECT().Update(ctx, origin).Return(nil)

	err := handler.Update(ctx, origin)
	assert.NoError(t, err)
}

func TestOriginHandlerGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockOriginStore(ctrl)
	handler := service.NewOriginHandler(mockStore)

	ctx := context.Background()
	id := uuid.New()
	expectedOrigin := service.Origin{
		ID:      id,
		Address: "http://localhost:9999",
	}

	mockStore.EXPECT().Get(ctx, id).Return(expectedOrigin, nil)

	origin, err := handler.Get(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedOrigin, origin)
}
