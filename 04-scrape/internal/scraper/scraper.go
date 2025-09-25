package scraper

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/zeuge/hw-go/04-scrape/internal/config"
)

type Scraper struct {
	cfg *config.ScraperConfig
}

func New(cfg *config.ScraperConfig) *Scraper {
	return &Scraper{
		cfg: cfg,
	}
}

func (s *Scraper) Run(ctx context.Context, chanUrls <-chan string) <-chan Result {
	chanResults := make(chan Result)

	var wg sync.WaitGroup

	for range s.cfg.Limit {
		wg.Go(func() {
			for url := range chanUrls {
				chanResults <- s.scrape(ctx, url)
			}
		})
	}

	go func() {
		wg.Wait()
		close(chanResults)
	}()

	return chanResults
}

func (s *Scraper) scrape(ctx context.Context, url string) Result {
	slog.Info("scraping", "url", url)

	ctx, cancel := context.WithTimeout(ctx, s.cfg.Timeout)
	defer cancel()

	client := &http.Client{}

	var (
		resp *http.Response
		err  error
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return NewResult(0, url, "", "")
	}

	for range s.cfg.Retries {
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()

			break
		}

		time.Sleep(s.cfg.Sleep)
	}

	if err != nil || resp == nil {
		slog.Error("scraping", "url", url, "error", err)

		return NewResult(0, url, "", "")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		slog.Error("parsing", "url", url, "error", err)

		return NewResult(resp.StatusCode, url, "", "")
	}

	title := doc.Find("title").First().Text()
	description, _ := doc.Find(`meta[name="description"]`).Attr("content")

	return NewResult(resp.StatusCode, url, title, description)
}
