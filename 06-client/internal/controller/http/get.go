package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

func (c *Controller) Get(ctx context.Context, cmd entity.Command) error {
	id, _ := cmd.Data["uuid"].(string)

	resp, err := c.requestWithContext(ctx, http.MethodGet, "/users/"+id, http.NoBody)
	if err != nil {
		return fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return entity.ErrUnexpectedHTTPStatus
	}

	var user entity.User

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	err = c.uc.GetUser(ctx, &user)
	if err != nil {
		return fmt.Errorf("c.uc.GetUser: %w", err)
	}

	return nil
}
