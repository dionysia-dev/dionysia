package redistore

import (
	"context"

	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type OriginStore struct {
	client *redis.Client
	ttl    int
}

func NewOriginStore(client *redis.Client, cfg *config.Config) *OriginStore {
	return &OriginStore{
		client: client,
		ttl:    cfg.OriginTTL,
	}
}

func (s *OriginStore) Update(ctx context.Context, origin service.Origin) error {
	var updateOrigin = redis.NewScript(`
		local key = KEYS[1]
		local value = ARGV[1]
		local ttl = ARGV[2]
		local curr_value = redis.call("GET", key)
		if curr_value == value or curr_value == false then
			redis.call("SETEX", key, ttl, value)
			return 1
		end
		return 0
	`)

	_, err := updateOrigin.Run(
		ctx,
		s.client,
		[]string{origin.ID.String()}, origin.Address, s.ttl).
		Result()

	return err
}

func (s *OriginStore) Get(ctx context.Context, uuid uuid.UUID) (service.Origin, error) {
	address, err := s.client.Get(ctx, uuid.String()).Result()
	if err != nil {
		return service.Origin{}, err
	}

	return service.Origin{
		ID:      uuid,
		Address: address,
	}, nil
}
