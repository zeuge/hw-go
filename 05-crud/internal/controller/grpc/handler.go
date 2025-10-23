package grpc

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeuge/hw-go/05-crud/internal/entity"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	role := roleFromProto(in.GetRole())

	user, err := entity.NewUser(in.GetName(), in.GetEmail(), role)
	if err != nil {
		return nil, fmt.Errorf("entity.NewUser: %w", err)
	}

	err = s.uc.CreateUser(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("entity.NewUser: %w", err)
	}

	res := &pb.CreateResponse{
		User: userToProto(&user),
	}

	return res, nil
}

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id, err := entity.NewIDFromString(in.GetUuid())
	if err != nil {
		return nil, fmt.Errorf("entity.NewIDFromString: %w", err)
	}

	user, err := s.uc.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.uc.GetUser: %w", err)
	}

	res := &pb.GetResponse{
		User: userToProto(user),
	}

	return res, nil
}

func (s *Server) GetAll(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllResponse, error) {
	users, err := s.uc.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("s.uc.GetUsers: %w", err)
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

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	id, err := entity.NewIDFromString(in.GetUuid())
	if err != nil {
		return nil, fmt.Errorf("entity.NewIDFromString: %w", err)
	}

	err = s.uc.DeleteUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.uc.DeleteUser: %w", err)
	}

	return &emptypb.Empty{}, nil
}
