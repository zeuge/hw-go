package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/adapter/grpc"
	"github.com/zeuge/hw-go/05-crud/internal/adapter/http"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/client"
)

func Run(ctx context.Context, cfg *config.Config, commands []*entity.Command) error {
	httpAdapter, err := http.New(ctx, &cfg.HTTPClient)
	if err != nil {
		return fmt.Errorf("http.New: %w", err)
	}

	grpcAdapter, err := grpc.New(ctx, &cfg.GRPCClient)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	var uc *usecase.UserUseCase

	if cfg.App.UseGRPC {
		uc = usecase.New(grpcAdapter)

		slog.InfoContext(ctx, "Run grpc.")
	} else {
		uc = usecase.New(httpAdapter)

		slog.InfoContext(ctx, "Run http.")
	}

	handler := NewHandler(uc)
	handler.Handle(ctx, commands)

	err = grpcAdapter.Close()
	if err != nil {
		return fmt.Errorf("grpcRepo.Close: %w", err)
	}

	return nil
}
