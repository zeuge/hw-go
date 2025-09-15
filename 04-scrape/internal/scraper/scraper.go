package scraper

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/zeuge/hw-go/04-scrape/internal/config"
)

type Result struct {
	Date        time.Time
	StatusCode  int
	URL         string
	Title       string
	Description string
}

func NewResult(statusCode int, url, title, description string) Result {
	return Result{
		Date:        time.Now(),
		StatusCode:  statusCode,
		URL:         url,
		Title:       title,
		Description: description,
	}
}

func (r *Result) ToSlice() []string {
	return []string{r.Date.Format(time.RFC3339), r.URL, strconv.Itoa(r.StatusCode), r.Title, r.Description}
}

func Run(chanUrls <-chan string, cfg *config.ScraperConfig) <-chan Result {
	chanResults := make(chan Result)

	var wg sync.WaitGroup

	for range cfg.Limit {
		wg.Go(func() {
			for url := range chanUrls {
				chanResults <- scrape(url, cfg)
			}
		})
	}

	go func() {
		wg.Wait()
		close(chanResults)
	}()

	return chanResults
}

func scrape(url string, cfg *config.ScraperConfig) Result {
	slog.Info("scraping", "url", url)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
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

	for range cfg.Retries {
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()

			break
		}

		time.Sleep(cfg.Sleep)
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
