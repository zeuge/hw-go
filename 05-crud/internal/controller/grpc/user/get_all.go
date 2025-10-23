package user

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

func (s *Server) GetAll(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllResponse, error) {
	users, err := s.uc.GetUsers(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.GetUsers", "error", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	usersProto := make([]*pb.User, 0, len(users))
	for _, user := range users {
		usersProto = append(usersProto, userToProto(&user))
	}

	res := &pb.GetAllResponse{
		Users: usersProto,
	}

	return res, nil
}
