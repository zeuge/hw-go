package grpc

import (
	"context"
	"fmt"

	"github.com/zeuge/hw-go/06-client/internal/entity"
	pb "github.com/zeuge/hw-go/06-client/pkg/client/grpc/user"
)

func (c *Controller) Delete(ctx context.Context, cmd entity.Command) error {
	id, _ := cmd.Data["uuid"].(string)

	req := pb.DeleteRequest{
		Uuid: id,
	}

	_, err := c.client.Delete(ctx, &req)
	if err != nil {
		return fmt.Errorf("c.client.Delete: %w", err)
	}

	err = c.uc.DeleteUser(ctx)
	if err != nil {
		return fmt.Errorf("c.uc.DeleteUser: %w", err)
	}

	return nil
}
