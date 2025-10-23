package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/zeuge/hw-go/06-client/config"
	"github.com/zeuge/hw-go/06-client/internal/usecase"
)

type Controller struct {
	name      string
	enabled   bool
	serverURL string
	uc        *usecase.UserUseCase
}

func New(cfg *config.HTTPConfig, uc *usecase.UserUseCase) *Controller {
	return &Controller{
		name:      "HTTP",
		serverURL: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		enabled:   cfg.Enabled,
		uc:        uc,
	}
}

func (c *Controller) Enabled() bool {
	return c.enabled
}

func (c *Controller) Name() string {
	return c.name
}

func (c *Controller) requestWithContext(
	ctx context.Context, method, path string, body io.Reader,
) (*http.Response, error) {
	url, err := url.JoinPath("http://", c.serverURL, path)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}

	return resp, nil
}
