package scraper

import (
	"strconv"
	"time"
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
