package grpc

import (
	"context"
	"fmt"

	"github.com/zeuge/hw-go/06-client/internal/entity"
	pb "github.com/zeuge/hw-go/06-client/pkg/client/grpc/user"
)

func (c *Controller) Get(ctx context.Context, cmd entity.Command) error {
	id, _ := cmd.Data["uuid"].(string)

	req := pb.GetRequest{
		Uuid: id,
	}

	resp, err := c.client.Get(ctx, &req)
	if err != nil {
		return fmt.Errorf("c.client.Get: %w", err)
	}

	protoUser := resp.GetUser()

	user, err := userFromProto(protoUser)
	if err != nil {
		return fmt.Errorf("userFromProto: %w", err)
	}

	err = c.uc.GetUser(ctx, user)
	if err != nil {
		return fmt.Errorf("c.uc.Get: %w", err)
	}

	return nil
}
