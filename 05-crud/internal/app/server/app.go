package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/controller/grpc"
	"github.com/zeuge/hw-go/05-crud/internal/controller/http"
	"github.com/zeuge/hw-go/05-crud/internal/repository/nats"
	"github.com/zeuge/hw-go/05-crud/internal/repository/pg"
	"github.com/zeuge/hw-go/05-crud/internal/repository/redis"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/server"
)

func Run(ctx context.Context, cfg *config.Config) error {
	repo, err := pg.New(ctx, &cfg.Pg)
	if err != nil {
		return fmt.Errorf("pg.New: %w", err)
	}
	defer repo.Close()

	cache := redis.New(ctx, &cfg.Redis)
	defer cache.Close()

	notify, err := nats.New(&cfg.NATS)
	if err != nil {
		return fmt.Errorf("nats.New: %w", err)
	}
	defer notify.Close()

	uc := usecase.New(repo, cache, notify)

	httpController := http.New(&cfg.HTTPServer, uc)

	go func() {
		err := httpController.Start()
		if err != nil {
			slog.ErrorContext(ctx, "httpController.Start", "error", err)
		}
	}()

	grpcController := grpc.New(&cfg.GRPCServer, uc)

	go func() {
		err := grpcController.Start(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "grpcController.Start", "error", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(ctx, cfg.App.GracefulShutdownTimeout)
	defer cancel()

	err = httpController.Stop(ctx)
	if err != nil {
		return fmt.Errorf("httpController.Stop: %w", err)
	}

	err = grpcController.Stop(ctx)
	if err != nil {
		return fmt.Errorf("grpcController.Stop: %w", err)
	}

	return nil
}
