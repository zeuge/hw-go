package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
)

type UserUseCase struct {
	repo UserRepository
}

func New(repo UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, dto dto.CreateUser) (*entity.User, error) {
	slog.InfoContext(ctx, "Start usecase.CreateUser")

	user, err := u.repo.Create(ctx, dto)
	if err != nil {
		return nil, fmt.Errorf("u.repo.Save: %w", err)
	}

	slog.InfoContext(ctx, "Finish usecase.CreateUser", "id", user.ID)

	return user, nil
}

func (u *UserUseCase) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	slog.InfoContext(ctx, "Start usecase.GetUser")

	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("u.repo.FindByID: %w", err)
	}

	slog.InfoContext(ctx, "Finish usecase.GetUser", "id", id)

	return user, nil
}

func (u *UserUseCase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	slog.InfoContext(ctx, "Start usecase.GetUsers")

	users, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("u.repo.FindAll: %w", err)
	}

	ids := make([]string, 0, len(users))

	for _, user := range users {
		ids = append(ids, user.ID.String())
	}

	slog.InfoContext(ctx, "Finish usecase.GetUsers", "ids", ids)

	return users, nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	slog.InfoContext(ctx, "Start usecase.DeleteUser")

	err := u.repo.DeleteByID(ctx, id)
	if err != nil {
		return fmt.Errorf("u.repo.DeleteByID: %w", err)
	}

	slog.InfoContext(ctx, "Finish usecase.DeleteUser", "id", id)

	return nil
}
