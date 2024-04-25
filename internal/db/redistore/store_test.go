//go:build integration

package redistore_test

import (
	"context"
	"testing"

	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/dionysia-dev/dionysia/internal/db/redistore"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestOriginStoreUpdate(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp"),
	}

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	assert.NoError(t, err, "could not start redis container")

	defer func() {
		assert.NoError(t, redisContainer.Terminate(ctx), "could not stop redis")
	}()

	endpoint, err := redisContainer.Endpoint(ctx, "")
	assert.NoError(t, err, "could not get redis endpoint")

	cfg := &config.Config{
		RedisAddr: endpoint,
		OriginTTL: 10,
	}

	client := redistore.NewClient(cfg)
	_, err = client.Ping(ctx).Result()
	assert.NoError(t, err, "could not ping redis")

	origin := service.Origin{
		ID:      uuid.New(),
		Address: "http://localhost:8080",
	}

	store := redistore.NewOriginStore(client, cfg)

	err = store.Update(ctx, origin)
	assert.NoError(t, err, "could not update origin")

	retrieved, err := store.Get(ctx, origin.ID)
	assert.NoError(t, err, "could not get origin")
	assert.Equal(t, origin.Address, retrieved.Address, "retrieved origin address does not match")

	// Update the origin with a different address
	// the address should not change because the TTL keeps it active
	newOrigin := service.Origin{
		ID:      uuid.New(),
		Address: "http://localhost:8081", // different address
	}

	err = store.Update(ctx, newOrigin)
	assert.NoError(t, err, "could not update origin")

	retrieved, err = store.Get(ctx, origin.ID)
	assert.NoError(t, err, "could not get origin")
	assert.Equal(t, origin.Address, retrieved.Address, "retrieved origin does not match the original one")
}
