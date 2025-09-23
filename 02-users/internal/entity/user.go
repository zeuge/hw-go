package entity

import (
	"fmt"
	"time"
)

type User struct {
	ID        ID
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
}

func NewUser(name, email string, role Role) (User, error) {
	id, err := NewID()
	if err != nil {
		return User{}, fmt.Errorf("NewID: %w", err)
	}

	user := User{
		ID:        id,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}

	return user, nil
}
