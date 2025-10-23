package user

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id, err := entity.NewIDFromString(in.GetUuid())
	if err != nil {
		slog.ErrorContext(ctx, "entity.NewIDFromString", "error", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.uc.GetUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.GetUser", "error", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pb.GetResponse{
		User: userToProto(user),
	}

	return res, nil
}
