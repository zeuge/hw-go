package app

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/zeuge/hw-go/06-client/config"
	"github.com/zeuge/hw-go/06-client/internal/controller"
	"github.com/zeuge/hw-go/06-client/internal/controller/grpc"
	"github.com/zeuge/hw-go/06-client/internal/controller/http"
	"github.com/zeuge/hw-go/06-client/internal/entity"
	"github.com/zeuge/hw-go/06-client/internal/usecase"
)

func Run(ctx context.Context, cfg *config.Config, commands []entity.Command) error {
	ch := make(chan entity.Command)

	go enqueueCommands(ctx, commands, ch)

	uc := usecase.New()
	httpController := http.New(&cfg.HTTP, uc)

	grpcController, err := grpc.New(&cfg.GRPC, uc)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	httpRouter := controller.NewRouter(httpController)
	grpcRouter := controller.NewRouter(grpcController)

	var wg sync.WaitGroup

	wg.Add(1)

	go worker(ctx, ch, &wg, httpRouter, grpcRouter)

	wg.Wait()

	err = grpcController.Stop()
	if err != nil {
		return fmt.Errorf("grpcController.Stop: %w", err)
	}

	return nil
}

func enqueueCommands(ctx context.Context, commands []entity.Command, ch chan<- entity.Command) {
	defer close(ch)

	for _, cmd := range commands {
		select {
		case <-ctx.Done():
			return
		case ch <- cmd:
		}
	}
}

func worker(
	ctx context.Context, ch <-chan entity.Command, wg *sync.WaitGroup, httpRouter, grpcRouter *controller.Router,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case cmd, ok := <-ch:
			if !ok {
				return
			}

			err := httpRouter.Handle(ctx, cmd)
			if err != nil {
				slog.ErrorContext(ctx, "httpRouter.Handle", "error", err)
			}

			err = grpcRouter.Handle(ctx, cmd)
			if err != nil {
				slog.ErrorContext(ctx, "grpcRouter.Handle", "error", err)
			}
		}
	}
}
