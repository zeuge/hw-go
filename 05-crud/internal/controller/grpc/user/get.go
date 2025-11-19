package user

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
)

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id, err := uuid.Parse(in.GetUuid())
	if err != nil {
		slog.ErrorContext(ctx, "uuid.Parse", "error", err)

		return nil, status.Error(codes.InvalidArgument, err.Error()) //nolint:wrapcheck
	}

	user, err := s.uc.GetUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.GetUser", "error", err)

		return nil, status.Error(codes.Internal, err.Error()) //nolint:wrapcheck
	}

	res := &pb.GetResponse{
		User: UserToProto(user),
	}

	return res, nil
}
