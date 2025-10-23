package user

import (
	"log/slog"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

func roleFromProto(role pb.Role) entity.Role {
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

func roleToProto(role entity.Role) pb.Role {
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

func userToProto(user *entity.User) *pb.User {
	proto := &pb.User{
		Uuid:      user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      roleToProto(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return proto
}
