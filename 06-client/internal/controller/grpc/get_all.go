package grpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

func (c *Controller) GetAll(ctx context.Context, cmd entity.Command) error {
	resp, err := c.client.GetAll(ctx, &empty.Empty{})
	if err != nil {
		return fmt.Errorf("c.client.GetUsers: %w", err)
	}

	protoUsers := resp.GetUsers()

	users := make([]*entity.User, 0, len(protoUsers))

	for _, protoUser := range protoUsers {
		user, err := userFromProto(protoUser)
		if err != nil {
			return fmt.Errorf("userFromProto: %w", err)
		}

		users = append(users, user)
	}

	err = c.uc.GetAllUsers(ctx, users)
	if err != nil {
		return fmt.Errorf("c.uc.GetAllUsers: %w", err)
	}

	return nil
}
