package grpc

import (
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/06-client/internal/entity"
	pb "github.com/zeuge/hw-go/06-client/pkg/client/grpc/user"
)

func userFromProto(proto *pb.User) (*entity.User, error) {
	id, err := entity.NewIDFromString(proto.GetUuid())
	if err != nil {
		return nil, fmt.Errorf("entity.NewIDFromString: %w", err)
	}

	user := entity.User{
		ID:        id,
		Name:      proto.GetName(),
		Email:     proto.GetEmail(),
		Role:      roleFromProto(proto.GetRole()),
		CreatedAt: proto.GetCreatedAt().AsTime(),
		UpdatedAt: proto.GetUpdatedAt().AsTime(),
	}

	return &user, nil
}

func roleToProto(role string) pb.Role {
	switch role {
	case entity.GuestRole:
		return pb.Role_GUEST
	case entity.AdminRole:
		return pb.Role_ADMIN
	case entity.UserRole:
		return pb.Role_USER
	default:
		slog.Warn("roleToProto: unknown role, use Guest", "role", role)

		return pb.Role_GUEST
	}
}

func roleFromProto(role pb.Role) string {
	switch role {
	case pb.Role_GUEST:
		return entity.GuestRole
	case pb.Role_ADMIN:
		return entity.AdminRole
	case pb.Role_USER:
		return entity.UserRole
	default:
		slog.Warn("roleFromProto: unknown role, use Guest", "role", pb.Role_name[int32(role)])

		return entity.GuestRole
	}
}
