package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/zeuge/hw-go/05-crud/api/pb/user"
	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/controller/grpc/user"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/server"
)

type Controller struct {
	server *grpc.Server
	health *health.Server
	port   string
}

func New(cfg *config.GRPCServerConfig, uc *usecase.UserUseCase) *Controller {
	server := grpc.NewServer()

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)

	pb.RegisterUserServiceServer(server, user.New(uc))

	controller := Controller{
		server: server,
		health: healthServer,
		port:   strconv.Itoa(cfg.Port),
	}

	return &controller
}

func (c *Controller) Start() error {
	lis, err := net.Listen("tcp", ":"+c.port) //nolint:noctx
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	slog.Info("🚀 GRPC server listening at " + lis.Addr().String())

	c.health.SetServingStatus(c.port, healthpb.HealthCheckResponse_SERVING)

	err = c.server.Serve(lis)
	if err != nil {
		return fmt.Errorf("c.server.Serve: %w", err)
	}

	return nil
}

func (c *Controller) Stop(ctx context.Context) error {
	c.health.SetServingStatus(c.port, healthpb.HealthCheckResponse_NOT_SERVING)

	stopped := make(chan struct{})

	go func() {
		c.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		c.server.Stop()

		slog.Info("📴 GRPC server force-stopped")
	case <-stopped:
		slog.Info("📴 GRPC server stopped gracefully")
	}

	return nil
}
