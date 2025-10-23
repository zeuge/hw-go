package usecase

import (
	"context"
	"log/slog"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

type UserUseCase struct{}

func New() *UserUseCase {
	return &UserUseCase{}
}

func (u *UserUseCase) GetUser(ctx context.Context, user *entity.User) error {
	slog.InfoContext(ctx, "usecase.GetUser", "id", user.ID)

	return nil
}

func (u *UserUseCase) GetAllUsers(ctx context.Context, users []*entity.User) error {
	ids := make([]string, 0, len(users))

	for _, user := range users {
		ids = append(ids, user.ID.String())
	}

	slog.InfoContext(ctx, "usecase.GetAllUsers", "ids", ids)

	return nil
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	slog.InfoContext(ctx, "usecase.CreateUser", "id", user.ID)

	return nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context) error {
	slog.InfoContext(ctx, "usecase.DeleteUser")

	return nil
}
