package memory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeuge/hw-go/02-users/internal/entity"
	"github.com/zeuge/hw-go/02-users/internal/repository/memory"
)

func newUser() entity.User {
	user := entity.NewUser("John Doe", "john.doe@example.com", entity.AdminRole)

	return user
}

func TestUserRepository_Save(t *testing.T) {
	repo := memory.NewUserRepository()
	user := newUser()

	err := repo.Save(user)
	assert.NoError(t, err)

	foundUser, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)
}

func TestUserRepository_FindByID(t *testing.T) {
	repo := memory.NewUserRepository()
	user := newUser()
	repo.Save(user)

	foundUser, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, foundUser)

	foundUser, err = repo.FindByID(entity.NewID())
	assert.Error(t, err)
}

func TestUserRepository_FindAll(t *testing.T) {
	repo := memory.NewUserRepository()
	user := newUser()
	repo.Save(user)

	users := repo.FindAll()
	assert.Equal(t, 1, len(users))
	assert.Equal(t, user, users[0])
}

func TestUserRepository_DeleteByID(t *testing.T) {
	repo := memory.NewUserRepository()
	user := newUser()
	repo.Save(user)

	err := repo.DeleteByID(user.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(user.ID)
	assert.Error(t, err)

	err = repo.DeleteByID(user.ID)
	assert.NoError(t, err)
}

func TestUserRepository_FindByRole(t *testing.T) {
	repo := memory.NewUserRepository()
	user := newUser()
	repo.Save(user)

	users := repo.FindByRole(entity.AdminRole)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, user, users[0])

	users = repo.FindByRole(entity.UserRole)
	assert.Equal(t, 0, len(users))
}
