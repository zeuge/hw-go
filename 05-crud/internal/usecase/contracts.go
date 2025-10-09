package usecase

import (
	"context"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

type (
	UserRepository interface {
		Ping(ctx context.Context) error
		Save(ctx context.Context, user *entity.User) error
		FindByID(ctx context.Context, id entity.ID) (*entity.User, error)
		FindAll(ctx context.Context) ([]entity.User, error)
		DeleteByID(ctx context.Context, id entity.ID) error
	}

	UserCacheRepository interface {
		Set(ctx context.Context, user *entity.User) error
		Get(ctx context.Context, id entity.ID) (*entity.User, error)
		Del(ctx context.Context, id entity.ID) error
	}

	NotificationRepository interface {
		Publish(subject string, message string) error
	}
)
