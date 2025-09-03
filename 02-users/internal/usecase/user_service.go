package usecase

import (
	"github.com/zeuge/hw-go/02-users/internal/entity"
	"github.com/zeuge/hw-go/02-users/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s UserService) CreateUser(name, email string, role entity.Role) (entity.User, error) {
	user := entity.NewUser(name, email, role)
	err := s.repo.Save(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s UserService) GetUser(id entity.ID) (entity.User, error) {
	return s.repo.FindByID(id)
}

func (s UserService) ListUsers() []entity.User {
	return s.repo.FindAll()
}

func (s UserService) RemoveUser(id entity.ID) error {
	return s.repo.DeleteByID(id)
}

func (s UserService) ListUsersByRole(role entity.Role) []entity.User {
	return s.repo.FindByRole(role)
}
