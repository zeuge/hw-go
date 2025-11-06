package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/zeuge/hw-go/04-scrape/internal/config"
	"github.com/zeuge/hw-go/04-scrape/internal/scraper"
)

func ReadLines(cfg *config.FileConfig) (<-chan string, <-chan error) {
	chanLines := make(chan string)
	chanErr := make(chan error, 1)

	go func() {
		defer close(chanLines)
		defer close(chanErr)

		file, err := os.Open(cfg.Input)
		if err != nil {
			chanErr <- fmt.Errorf("os.Open: %w", err)

			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			chanLines <- scanner.Text()
		}

		err = scanner.Err()
		if err != nil {
			chanErr <- fmt.Errorf("scanner.Err: %w", err)
		}
	}()

	return chanLines, chanErr
}

func WriteResults(cfg *config.FileConfig, chanResults <-chan scraper.Result) error {
	file, err := os.Create(cfg.Output)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	defer writer.Flush()

	headers := []string{"date", "url", "status_code", "title", "description"}

	err = writer.Write(headers)
	if err != nil {
		return fmt.Errorf("writer.Write: %w", err)
	}

	for r := range chanResults {
		err := writer.Write(r.ToSlice())
		if err != nil {
			return fmt.Errorf("writer.Write: %w", err)
		}
	}

	return nil
}
