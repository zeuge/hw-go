package user

import (
	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/server"
)

type Server struct {
	pb.UnimplementedUserServiceServer

	uc *usecase.UserUseCase
}

func New(uc *usecase.UserUseCase) *Server {
	s := &Server{
		uc: uc,
	}

	return s
}
