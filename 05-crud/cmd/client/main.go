package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/zeuge/hw-go/05-crud/config"
	app "github.com/zeuge/hw-go/05-crud/internal/app/client"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
)

//go:embed commands.json
var b []byte

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.ReadConfig()
	if err != nil {
		slog.ErrorContext(ctx, "config.ReadConfig", "error", err)

		return
	}

	var commands []*entity.Command

	err = json.Unmarshal(b, &commands)
	if err != nil {
		slog.ErrorContext(ctx, "json.Unmarshal", "error", err)

		return
	}

	err = app.Run(ctx, cfg, commands)
	if err != nil {
		slog.ErrorContext(ctx, "app.Run", "error", err)
	}
}
