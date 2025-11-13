package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/controller/grpc/user"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
)

type GRPCUserAdapter struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

func New(ctx context.Context, cfg *config.GRPCClientConfig) (*GRPCUserAdapter, error) {
	target := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := pb.NewUserServiceClient(conn)

	repo := &GRPCUserAdapter{
		conn:   conn,
		client: client,
	}

	return repo, nil
}

func (a *GRPCUserAdapter) Close() error {
	err := a.conn.Close()
	if err != nil {
		return fmt.Errorf("r.conn.Close: %w", err)
	}

	return nil
}

func (a *GRPCUserAdapter) Create(ctx context.Context, dto dto.CreateUser) (*entity.User, error) {
	req := pb.CreateRequest{
		Name:  dto.Name,
		Email: dto.Email,
		Role:  user.RoleToProto(dto.Role),
	}

	resp, err := a.client.Create(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("c.client.Create: %w", err)
	}

	protoUser := resp.GetUser()

	user, err := user.UserFromProto(protoUser)
	if err != nil {
		return nil, fmt.Errorf("user.UserFromProto: %w", err)
	}

	return user, nil
}

func (a *GRPCUserAdapter) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	req := pb.GetRequest{
		Uuid: id.String(),
	}

	resp, err := a.client.Get(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("r.client.Get: %w", err)
	}

	protoUser := resp.GetUser()

	user, err := user.UserFromProto(protoUser)
	if err != nil {
		return nil, fmt.Errorf("user.UserFromProto: %w", err)
	}

	return user, nil
}

func (a *GRPCUserAdapter) FindAll(ctx context.Context) ([]*entity.User, error) {
	resp, err := a.client.GetAll(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("r.client.GetUsers: %w", err)
	}

	protoUsers := resp.GetUsers()

	users := make([]*entity.User, 0, len(protoUsers))

	for _, protoUser := range protoUsers {
		user, err := user.UserFromProto(protoUser)
		if err != nil {
			return nil, fmt.Errorf("user.UserFromProto: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (a *GRPCUserAdapter) DeleteByID(ctx context.Context, id uuid.UUID) error {
	req := pb.DeleteRequest{
		Uuid: id.String(),
	}

	_, err := a.client.Delete(ctx, &req)
	if err != nil {
		return fmt.Errorf("r.client.Delete: %w", err)
	}

	return nil
}
