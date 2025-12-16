package user

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

func RoleFromProto(role pb.Role) string {
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

func RoleToProto(role string) pb.Role {
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

func UserFromProto(proto *pb.User) (*entity.User, error) {
	id, err := uuid.Parse(proto.GetUuid())
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}

	user := entity.User{
		ID:        id,
		Name:      proto.GetName(),
		Email:     proto.GetEmail(),
		Role:      RoleFromProto(proto.GetRole()),
		CreatedAt: proto.GetCreatedAt().AsTime(),
		UpdatedAt: proto.GetUpdatedAt().AsTime(),
	}

	return &user, nil
}

func UserToProto(user *entity.User) *pb.User {
	proto := &pb.User{
		Uuid:      user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      RoleToProto(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return proto
}
