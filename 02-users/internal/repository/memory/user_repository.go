package memory

import (
	"github.com/zeuge/hw-go/02-users/internal/entity"
)

type InMemoryUserRepo struct {
	users map[entity.ID]*entity.User
}

func NewUserRepository() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[entity.ID]*entity.User),
	}
}

func (r *InMemoryUserRepo) Save(user entity.User) error {
	r.users[user.ID] = &user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id entity.ID) (entity.User, error) {
	user, ok := r.users[id]
	if !ok {
		return entity.User{}, entity.ErrNotFound
	}
	return *user, nil
}

func (r *InMemoryUserRepo) FindAll() []entity.User {
	users := make([]entity.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}

	return users
}

func (r *InMemoryUserRepo) DeleteByID(id entity.ID) error {
	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepo) FindByRole(role entity.Role) []entity.User {
	users := make([]entity.User, 0, len(r.users))
	for _, user := range r.users {
		if user.Role == role {
			users = append(users, *user)
		}
	}
	return users
}
