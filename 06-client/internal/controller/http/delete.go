package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/zeuge/hw-go/06-client/internal/entity"
)

func (c *Controller) Delete(ctx context.Context, cmd entity.Command) error {
	id, _ := cmd.Data["uuid"].(string)

	resp, err := c.requestWithContext(ctx, http.MethodDelete, "/users/"+id, http.NoBody)
	if err != nil {
		return fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		slog.ErrorContext(ctx, "unexpected status", "code", resp.StatusCode, "status", resp.Status)

		return entity.ErrUnexpectedHTTPStatus
	}

	err = c.uc.DeleteUser(ctx)
	if err != nil {
		return fmt.Errorf("c.uc.DeleteUser: %w", err)
	}

	return nil
}
