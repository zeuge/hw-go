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

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/controller/grpc/user"
	"github.com/zeuge/hw-go/05-crud/internal/usecase"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

type Controller struct {
	server *grpc.Server
	health *health.Server
	port   string
}

func New(cfg *config.GRPCConfig, uc *usecase.UserUseCase) *Controller {
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

func (c *Controller) Start(ctx context.Context) error {
	listenConfig := net.ListenConfig{}

	listener, err := listenConfig.Listen(ctx, "tcp", ":"+c.port)
	if err != nil {
		return fmt.Errorf("listenConfig.Listen: %w", err)
	}

	slog.Info("🚀 GRPC server listening at " + listener.Addr().String())

	c.health.SetServingStatus(c.port, healthpb.HealthCheckResponse_SERVING)

	err = c.server.Serve(listener)
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
