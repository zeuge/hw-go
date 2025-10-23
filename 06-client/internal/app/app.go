package app

import (
	"context"

	"github.com/zeuge/hw-go/06-client/config"
	"github.com/zeuge/hw-go/06-client/internal/controller/grpc"
	"github.com/zeuge/hw-go/06-client/internal/controller/http"
	"github.com/zeuge/hw-go/06-client/internal/entity"
	"github.com/zeuge/hw-go/06-client/internal/usecase"
)

func Run(ctx context.Context, cfg *config.Config, commands []entity.Command) error {
	done := make(chan struct{})

	uc := usecase.New()

	httpController := http.New(&cfg.HTTP, uc)
	grpcController := grpc.New(&cfg.GRPC, uc)

	go func() {
		for _, cmd := range commands {
			httpController.Run(ctx, cmd)
			grpcController.Run(ctx, cmd)
		}

		close(done)
	}()

	select {
	case <-ctx.Done():
	case <-done:
	}

	// ctx, cancel := context.WithTimeout(ctx, cfg.App.GracefulShutdownTimeout)
	// defer cancel()

	// err = controller.Stop(ctx)
	// if err != nil {
	// 	return fmt.Errorf("controller.Stop: %w", err)
	// }

	return nil
}
