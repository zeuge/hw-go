package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/app/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.ReadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "config.ReadConfig", "error", err)

		return
	}

	err = server.Run(ctx, cfg)
	if err != nil {
		slog.ErrorContext(ctx, "app.Run", "error", err)
	}
}
