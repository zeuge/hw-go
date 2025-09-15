package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/zeuge/hw-go/04-scrape/internal/app"
	"github.com/zeuge/hw-go/04-scrape/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("app.LoadConfig: %v", err))
		os.Exit(1)
	}

	err = app.Run(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("app.Run: %v", err))
		os.Exit(1)
	}
}
