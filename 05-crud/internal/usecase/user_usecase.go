package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

type UserUseCase struct {
	repo   UserRepository
	cache  UserCacheRepository
	notify NotificationRepository
}

func New(repo UserRepository, cache UserCacheRepository, notify NotificationRepository) *UserUseCase {
	return &UserUseCase{repo, cache, notify}
}

func (s *UserUseCase) DbCheck(ctx context.Context) error {
	err := s.repo.Ping(ctx)
	if err != nil {
		return fmt.Errorf("s.repo.Check: %w", err)
	}

	return nil
}

func (s *UserUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	err := s.notify.Publish("info", "CreateUser")
	if err != nil {
		slog.ErrorContext(ctx, "failed to publish notification", "action", "CreateUser", "error", err)
	}

	err = s.repo.Save(ctx, user)
	if err != nil {
		return fmt.Errorf("s.repo.Save: %w", err)
	}

	err = s.cache.Set(ctx, user)
	if err != nil {
		return fmt.Errorf("s.cache.Set: %w", err)
	}

	return nil
}

func (s *UserUseCase) GetUser(ctx context.Context, id entity.ID) (*entity.User, error) {
	err := s.notify.Publish("info", "GetUser")
	if err != nil {
		slog.ErrorContext(ctx, "failed to publish notification", "action", "GetUser", "error", err)
	}

	user, err := s.cache.Get(ctx, id)
	if err != nil {
		slog.WarnContext(ctx, "failed to get user from cache", "id", id)
	} else {
		return user, nil
	}

	user, err = s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.repo.FindByID: %w", err)
	}

	return user, nil
}

func (s *UserUseCase) GetUsers(ctx context.Context) ([]entity.User, error) {
	err := s.notify.Publish("info", "GetUsers")
	if err != nil {
		slog.ErrorContext(ctx, "failed to publish notification", "action", "GetUsers", "error", err)
	}

	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("s.repo.FindAll: %w", err)
	}

	return users, nil
}

func (s *UserUseCase) DeleteUser(ctx context.Context, id entity.ID) error {
	err := s.notify.Publish("info", "GetUsers")
	if err != nil {
		slog.ErrorContext(ctx, "failed to publish notification", "action", "DeleteUser", "error", err)
	}

	err = s.repo.DeleteByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.repo.DeleteByID: %w", err)
	}

	err = s.cache.Del(ctx, id)
	if err != nil {
		return fmt.Errorf("s.cache.Del: %w", err)
	}

	return nil
}
