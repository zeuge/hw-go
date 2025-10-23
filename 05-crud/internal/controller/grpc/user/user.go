package user

import (
	"github.com/zeuge/hw-go/05-crud/internal/usecase"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
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
