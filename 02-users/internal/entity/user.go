package entity

import "time"

type User struct {
	ID        ID
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
}

func NewUser(name, email string, role Role) User {
	return User{
		ID:        NewID(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
}
