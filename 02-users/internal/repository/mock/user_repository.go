package mock

import (
	"log/slog"

	"github.com/zeuge/hw-go/02-users/internal/entity"
)

type MockUserRepo struct{}

func NewUserRepository() *MockUserRepo {
	return &MockUserRepo{}
}

func (r *MockUserRepo) Save(user entity.User) error {
	slog.Info(
		"call Save with params",
		slog.Group("user",
			slog.String("id", user.ID.String()),
			slog.String("name", user.Name),
			slog.String("email", user.Email),
			slog.String("role", user.Role.String()),
			slog.String("created_at", user.CreatedAt.String()),
		),
	)
	return nil
}

func (r *MockUserRepo) FindByID(id entity.ID) (entity.User, error) {
	slog.Info("call FindByID with params: ", slog.String("id", id.String()))
	return entity.User{}, nil
}

func (r *MockUserRepo) FindAll() []entity.User {
	slog.Info("call FindAll")
	return []entity.User{}
}

func (r *MockUserRepo) DeleteByID(id entity.ID) error {
	slog.Info("call DeleteByID with params: ", slog.String("id", id.String()))
	return nil
}

func (r *MockUserRepo) FindByRole(role entity.Role) []entity.User {
	slog.Info("call FindByRole with params: ", slog.String("role", role.String()))
	return []entity.User{}
}
