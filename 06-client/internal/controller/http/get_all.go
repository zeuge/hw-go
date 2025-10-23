package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

func (c *Controller) GetAll(ctx context.Context, cmd entity.Command) error {
	resp, err := c.requestWithContext(ctx, http.MethodGet, "/users", http.NoBody)
	if err != nil {
		return fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return entity.ErrUnexpectedHTTPStatus
	}

	var users []*entity.User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	err = c.uc.GetAllUsers(ctx, users)
	if err != nil {
		return fmt.Errorf("c.uc.GetAllUsers: %w", err)
	}

	return nil
}
