package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/uuid"

	"github.com/zeuge/hw-go/05-crud/config"
	"github.com/zeuge/hw-go/05-crud/internal/entity"
	"github.com/zeuge/hw-go/05-crud/internal/entity/dto"
)

type HTTPUserAdapter struct {
	serverURL string
}

func New(ctx context.Context, cfg *config.HTTPClientConfig) (*HTTPUserAdapter, error) {
	port := strconv.Itoa(cfg.Port)
	target := "http://" + net.JoinHostPort(cfg.Host, port)

	repo := &HTTPUserAdapter{
		serverURL: target,
	}

	return repo, nil
}

func (a *HTTPUserAdapter) Create(ctx context.Context, dto dto.CreateUser) (*entity.User, error) {
	data, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	resp, err := a.requestWithContext(ctx, http.MethodPost, "/users", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("r.requestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return nil, entity.ErrUnexpectedHTTPStatus
	}

	user := new(entity.User)

	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return user, nil
}

func (a *HTTPUserAdapter) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	resp, err := a.requestWithContext(ctx, http.MethodGet, "/users/"+id.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("r.requestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return nil, entity.ErrUnexpectedHTTPStatus
	}

	user := new(entity.User)

	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return user, nil
}

func (a *HTTPUserAdapter) FindAll(ctx context.Context) ([]*entity.User, error) {
	resp, err := a.requestWithContext(ctx, http.MethodGet, "/users", http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("r.requestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return nil, entity.ErrUnexpectedHTTPStatus
	}

	var users []*entity.User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return users, nil
}

func (a *HTTPUserAdapter) DeleteByID(ctx context.Context, id uuid.UUID) error {
	resp, err := a.requestWithContext(ctx, http.MethodDelete, "/users/"+id.String(), http.NoBody)
	if err != nil {
		return fmt.Errorf("r.requestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return entity.ErrUnexpectedHTTPStatus
	}

	return nil
}

func (a *HTTPUserAdapter) requestWithContext(
	ctx context.Context, method, path string, body io.Reader,
) (*http.Response, error) {
	url, err := url.JoinPath(a.serverURL, path)
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
