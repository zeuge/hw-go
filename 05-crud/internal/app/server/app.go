package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/adapter/nats"
	"github.com/zeuge/hw-go/05-crud/internal/controller/grpc"
	"github.com/zeuge/hw-go/05-crud/internal/controller/http"
	"github.com/zeuge/hw-go/05-crud/internal/repository/pg"
	"github.com/zeuge/hw-go/05-crud/internal/repository/redis"
	"github.com/zeuge/hw-go/05-crud/internal/tracing"
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

	tracer, err := tracing.New(ctx, &cfg.Tracing)
	if err != nil {
		return fmt.Errorf("tracing.New: %w", err)
	}

	go func() {
		err := httpController.Start()
		if err != nil {
			slog.ErrorContext(ctx, "httpController.Start", "error", err)
		}
	}()

	grpcController := grpc.New(&cfg.GRPCServer, uc)

	go func() {
		err := grpcController.Start()
		if err != nil {
			slog.ErrorContext(ctx, "grpcController.Start", "error", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.GracefulShutdownTimeout) //nolint:contextcheck
	defer cancel()

	err = httpController.Stop(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "httpController.Stop", "error", err)
	}

	err = grpcController.Stop(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "grpcController.Stop", "error", err)
	}

	err = tracer.Shutdown(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "grpcController.Stop", "error", err)
	}

	return nil
}
