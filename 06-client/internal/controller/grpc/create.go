package grpc

import (
	"context"
	"fmt"

	"github.com/zeuge/hw-go/06-client/internal/entity"
	pb "github.com/zeuge/hw-go/06-client/pkg/client/grpc/user"
)

func (c *Controller) Create(ctx context.Context, cmd entity.Command) error {
	name, _ := cmd.Data["name"].(string)
	email, _ := cmd.Data["email"].(string)
	role, _ := cmd.Data["role"].(string)
	roleProto := roleToProto(role)

	req := pb.CreateRequest{
		Name:  name,
		Email: email,
		Role:  roleProto,
	}

	resp, err := c.client.Create(ctx, &req)
	if err != nil {
		return fmt.Errorf("c.client.Create: %w", err)
	}

	protoUser := resp.GetUser()

	user, err := userFromProto(protoUser)
	if err != nil {
		return fmt.Errorf("userFromProto: %w", err)
	}

	err = c.uc.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("c.uc.CreateUser: %w", err)
	}

	return nil
}
