package user

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	role := RoleFromProto(in.GetRole())

	user, err := entity.NewUser(in.GetName(), in.GetEmail(), role)
	if err != nil {
		slog.ErrorContext(ctx, "entity.NewUser", "error", err)

		return nil, status.Error(codes.InvalidArgument, err.Error()) //nolint:wrapcheck
	}

	err = s.uc.CreateUser(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.CreateUser", "error", err)

		return nil, status.Error(codes.Internal, err.Error()) //nolint:wrapcheck
	}

	res := &pb.CreateResponse{
		User: UserToProto(user),
	}

	return res, nil
}
