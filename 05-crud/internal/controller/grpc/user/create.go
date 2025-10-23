package user

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	role := roleFromProto(in.GetRole())

	user, err := entity.NewUser(in.GetName(), in.GetEmail(), role)
	if err != nil {
		slog.ErrorContext(ctx, "entity.NewUser", "error", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.uc.CreateUser(ctx, &user)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.CreateUser", "error", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.CreateResponse{
		User: userToProto(&user),
	}

	return res, nil
}
