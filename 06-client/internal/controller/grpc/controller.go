package grpc

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/zeuge/hw-go/06-client/config"
	"github.com/zeuge/hw-go/06-client/internal/usecase"
	pb "github.com/zeuge/hw-go/06-client/pkg/client/grpc/user"
)

type Controller struct {
	conn    *grpc.ClientConn
	client  pb.UserServiceClient
	name    string
	enabled bool
	uc      *usecase.UserUseCase
}

func New(cfg *config.GRPCConfig, uc *usecase.UserUseCase) (*Controller, error) {
	target := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := pb.NewUserServiceClient(conn)

	controller := Controller{
		conn:    conn,
		client:  client,
		name:    "GRPC",
		enabled: cfg.Enabled,
		uc:      uc,
	}

	return &controller, nil
}

func (c *Controller) Stop() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("c.conn.Close: %w", err)
	}

	return nil
}

func (c *Controller) Enabled() bool {
	return c.enabled
}

func (c *Controller) Name() string {
	return c.name
}
