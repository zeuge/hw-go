package app

import (
	"fmt"
	"log/slog"

	"github.com/zeuge/hw-go/04-scrape/internal/config"
	"github.com/zeuge/hw-go/04-scrape/internal/file"
	"github.com/zeuge/hw-go/04-scrape/internal/scraper"
)

func Run(cfg *config.Config) error {
	slog.Info("run app", "config", cfg)

	chanUrls, chanErr := file.ReadLines(&cfg.File)

	chanResults := scraper.Run(chanUrls, &cfg.Scraper)

	err := file.WriteResults(&cfg.File, chanResults)
	if err != nil {
		return fmt.Errorf("file.WriteResults: %w", err)
	}

	err = <-chanErr
	if err != nil {
		return fmt.Errorf("file.ReadLines: %w", err)
	}

	return nil
}
