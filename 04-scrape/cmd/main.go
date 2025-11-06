package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/zeuge/hw-go/04-scrape/internal/app"
	"github.com/zeuge/hw-go/04-scrape/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("app.LoadConfig: %v", err))

		return
	}

	err = app.Run(ctx, cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("app.Run: %v", err))

		return
	}
}
