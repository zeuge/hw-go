package main

import (
	"fmt"

	"github.com/zeuge/hw-go/02-users/internal/entity"
	"github.com/zeuge/hw-go/02-users/internal/repository/mock"
	"github.com/zeuge/hw-go/02-users/internal/usecase"
)

func main() {
	repo := mock.NewUserRepository()
	// repo := memory.NewUserRepository()
	svc := usecase.NewUserService(repo)

	user, err := svc.CreateUser("John Doe", "john.doe@example.com", entity.AdminRole)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	user, err = svc.GetUser(user.ID)
	if err != nil {
		fmt.Println("Error getting user:", err)
		return
	}
	fmt.Printf("GetUser result: %+v\n", user)

	users := svc.ListUsers()
	fmt.Printf("ListUsers result: %+v\n", users)

	users = svc.ListUsersByRole(entity.AdminRole)
	fmt.Printf("ListUsersByRole result: %+v\n", users)

	svc.RemoveUser(user.ID)
}
