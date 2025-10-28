package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/controller/client"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/repository/grpc"
	"github.com/zeuge/hw-go/05-crud/internal/repository/http"
	usecase "github.com/zeuge/hw-go/05-crud/internal/usecase/client"
)

func Run(ctx context.Context, cfg *config.Config, commands []*entity.Command) error {
	httpRepo, err := http.New(ctx, &cfg.HTTPClient)
	if err != nil {
		return fmt.Errorf("http.New: %w", err)
	}

	grpcRepo, err := grpc.New(ctx, &cfg.GRPCClient)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	var uc *usecase.UserUseCase

	if cfg.App.UseGRPC {
		uc = usecase.New(grpcRepo)

		slog.InfoContext(ctx, "Run grpc.")
	} else {
		uc = usecase.New(httpRepo)

		slog.InfoContext(ctx, "Run http.")
	}

	controller := client.New(uc, commands)

	ch := make(chan struct{})

	go func() {
		defer close(ch)

		controller.Start(ctx)
	}()

	select {
	case <-ctx.Done():
	case <-ch:
	}

	err = grpcRepo.Close()
	if err != nil {
		return fmt.Errorf("grpcRepo.Close: %w", err)
	}

	return nil
}
