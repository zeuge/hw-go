package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/usecase"
	pb "github.com/zeuge/hw-go/05-crud/pkg/server/grpc/user"
)

type Server struct {
	pb.UnimplementedUserServiceServer

	addr string
	grpc *grpc.Server
	uc   *usecase.UserUseCase
}

func NewServer(cfg *config.GRPCConfig, uc *usecase.UserUseCase) *Server {
	return &Server{
		addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		grpc: grpc.NewServer(),
		uc:   uc,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listenConfig := net.ListenConfig{}

	listener, err := listenConfig.Listen(ctx, "tcp", s.addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	pb.RegisterUserServiceServer(s.grpc, s)
	slog.Info("🚀 GRPC server listening at " + listener.Addr().String())

	err = s.grpc.Serve(listener)
	if err != nil {
		return fmt.Errorf("server.Serve: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	stopped := make(chan struct{})

	go func() {
		s.grpc.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		s.grpc.Stop()

		slog.Info("📴 GRPC server force-stopped")
	case <-stopped:
		slog.Info("📴 GRPC server stopped gracefully")
	}

	return nil
}
