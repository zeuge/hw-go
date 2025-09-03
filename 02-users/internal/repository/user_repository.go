package repository

import "github.com/zeuge/hw-go/02-users/internal/entity"

type UserRepository interface {
	Save(user entity.User) error
	FindByID(id entity.ID) (entity.User, error)
	FindAll() []entity.User
	DeleteByID(id entity.ID) error
	FindByRole(role entity.Role) []entity.User
}
