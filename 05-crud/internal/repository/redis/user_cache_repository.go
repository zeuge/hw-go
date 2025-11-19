package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

type UserCacheRepository struct {
	client     *redis.Client
	expiration time.Duration
}

func New(ctx context.Context, cfg *config.RedisConfig) *UserCacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	})

	cache := UserCacheRepository{
		client:     client,
		expiration: cfg.Expiration,
	}

	return &cache
}

func (r *UserCacheRepository) Close() error {
	err := r.client.Close()
	if err != nil {
		return fmt.Errorf("r.client.Close: %w", err)
	}

	return nil
}

func (r *UserCacheRepository) Set(ctx context.Context, user *entity.User) error {
	buf, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = r.client.Set(ctx, user.ID.String(), buf, r.expiration).Err()
	if err != nil {
		return fmt.Errorf("r.client.Set: %w", err)
	}

	return nil
}

func (r *UserCacheRepository) Get(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	buf, err := r.client.Get(ctx, id.String()).Bytes()
	if err != nil {
		return nil, fmt.Errorf("r.client.Get: %w", err)
	}

	var user entity.User

	err = json.Unmarshal(buf, &user)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return &user, nil
}

func (r *UserCacheRepository) Del(ctx context.Context, id uuid.UUID) error {
	err := r.client.Del(ctx, id.String()).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf(".client.Del: %w", err)
	}

	return nil
}
