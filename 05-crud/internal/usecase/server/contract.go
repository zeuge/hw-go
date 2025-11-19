package server

import (
	"context"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

type (
	UserRepository interface {
		Ping(ctx context.Context) error
		Save(ctx context.Context, user *entity.User) error
		FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
		FindAll(ctx context.Context) ([]*entity.User, error)
		DeleteByID(ctx context.Context, id uuid.UUID) error
	}

	UserCacheRepository interface {
		Set(ctx context.Context, user *entity.User) error
		Get(ctx context.Context, id uuid.UUID) (*entity.User, error)
		Del(ctx context.Context, id uuid.UUID) error
	}

	NotificationRepository interface {
		Publish(subject string, message string) error
	}
)
