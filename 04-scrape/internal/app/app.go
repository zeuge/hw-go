package app

import (
	"context"
	"fmt"

	"github.com/zeuge/hw-go/04-scrape/internal/config"
	"github.com/zeuge/hw-go/04-scrape/internal/file"
	"github.com/zeuge/hw-go/04-scrape/internal/scraper"
)

func Run(ctx context.Context, cfg *config.Config) error {
	chanUrls, chanErr := file.ReadLines(&cfg.File)

	s := scraper.New(&cfg.Scraper)
	chanResults := s.Run(ctx, chanUrls)

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
