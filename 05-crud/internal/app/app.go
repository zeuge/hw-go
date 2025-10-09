package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/controller/http"
	"github.com/zeuge/hw-go/05-crud/internal/repository/nats"
	"github.com/zeuge/hw-go/05-crud/internal/repository/pg"
	"github.com/zeuge/hw-go/05-crud/internal/repository/redis"
	"github.com/zeuge/hw-go/05-crud/internal/usecase"
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
	controller := http.New(&cfg.HTTP, uc)

	go func() {
		err := controller.Start()
		if err != nil {
			slog.ErrorContext(ctx, "controller.Start", "error", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(ctx, cfg.App.GracefulShutdownTimeout)
	defer cancel()

	err = controller.Stop(ctx)
	if err != nil {
		return fmt.Errorf("controller.Stop: %w", err)
	}

	return nil
}
