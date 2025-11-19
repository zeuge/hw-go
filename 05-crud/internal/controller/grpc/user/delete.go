package user

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
)

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(in.GetUuid())
	if err != nil {
		slog.ErrorContext(ctx, "uuid.Parse", "error", err)

		return nil, status.Error(codes.InvalidArgument, err.Error()) //nolint:wrapcheck
	}

	err = s.uc.DeleteUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.DeleteUser", "error", err)

		return nil, status.Error(codes.Internal, err.Error()) //nolint:wrapcheck
	}

	return &emptypb.Empty{}, nil
}
