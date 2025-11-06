package usecase

import (
	"fmt"

	"github.com/zeuge/hw-go/02-users/internal/entity"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string, role entity.Role) (entity.User, error) {
	user, err := entity.NewUser(name, email, role)
	if err != nil {
		return entity.User{}, fmt.Errorf("entity.NewUser: %w", err)
	}

	err = s.repo.Save(user)
	if err != nil {
		return entity.User{}, fmt.Errorf("s.repo.Save: %w", err)
	}
	return user, nil
}

func (s *UserService) GetUser(id entity.ID) (entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) ListUsers() []entity.User {
	return s.repo.FindAll()
}

func (s *UserService) RemoveUser(id entity.ID) error {
	return s.repo.DeleteByID(id)
}

func (s *UserService) ListUsersByRole(role entity.Role) []entity.User {
	return s.repo.FindByRole(role)
}
