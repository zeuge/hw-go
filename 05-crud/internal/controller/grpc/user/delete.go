package user

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id, err := entity.NewIDFromString(in.GetUuid())
	if err != nil {
		slog.ErrorContext(ctx, "entity.NewIDFromString", "error", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.uc.DeleteUser(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "s.uc.DeleteUser", "error", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
